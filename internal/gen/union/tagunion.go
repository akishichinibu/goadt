package union

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
)

type TagUnionGenerator struct {
	order  int
	naming unionNaming
}

func NewTagUnionGenerator(n int) *TagUnionGenerator {
	return &TagUnionGenerator{
		order:  n,
		naming: newUnionNaming(n),
	}
}

func (g *TagUnionGenerator) typeNameIds() (cs []j.Code) {
	for i := range g.order {
		cs = append(cs, j.Id(g.naming.typeName(i+1)))
	}
	return cs
}

func (g *TagUnionGenerator) typeNamesWithAny() (cs []j.Code) {
	for id := range g.typeNameIds() {
		cs = append(cs, j.Id(g.naming.typeName(id+1)).Any())
	}
	return cs
}

func (g *TagUnionGenerator) paramsWithSingleParamFunc() (cs []j.Code) {
	for i := 1; i <= g.order; i++ {
		cs = append(cs, j.Id(g.naming.value(i)).Func().Params(
			j.Id(g.naming.value(i)).Id(g.naming.typeName(i)),
		))
	}
	return cs
}

func (g *TagUnionGenerator) genStruct(f *j.File) {
	f.Type().Id(g.naming.structName).Types(g.typeNamesWithAny()...).StructFunc(func(s *j.Group) {
		s.Id(g.naming.kindMember).Uint8()
		for i := 1; i <= g.order; i++ {
			s.Id(g.naming.value(i)).Id(g.naming.typeName(i))
		}
	})
}

func (g *TagUnionGenerator) GenWhenMethod(f *j.File) {
	f.Func().
		Params(j.Id("u").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
		Id("When").
		Params(g.paramsWithSingleParamFunc()...).BlockFunc(func(b *j.Group) {
		b.Switch(j.Id("u").Dot(g.naming.kindMember)).BlockFunc(func(s *j.Group) {
			for i := 1; i <= g.order; i++ {
				s.Case(j.Lit(i)).Block(
					j.Id(g.naming.value(i)).Call(j.Id("u").Dot(g.naming.value(i))),
				)
			}
			s.Default().Block(j.Panic(j.Lit("unreachable")))
		})
	})
}

// Generate Get1 ~ GetN methods
func (g *TagUnionGenerator) GenGetMethods(f *j.File) {
	for i := 1; i <= g.order; i++ {
		f.Func().Params(j.Id("u").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
			Id(g.naming.getter(i)).Params().
			Params(j.Id(g.naming.typeName(i)), j.Bool()).
			Block(
				j.If(j.Id("u").Dot(g.naming.kindMember).Op("==").Lit(i)).Block(
					j.Return(j.Id("u").Dot(g.naming.value(i)), j.True()),
				),
				j.Return(j.Id("u").Dot(g.naming.value(i)), j.False()),
			)
	}
}

func (g *TagUnionGenerator) GenJSONMethods(f *j.File) {
	// MarshalJSON
	f.Func().Params(j.Id("u").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
		Id("MarshalJSON").
		Params().Params(j.Index().Byte(), j.Error()).
		BlockFunc(func(b *j.Group) {
			b.Switch(j.Id("u").Dot(g.naming.kindMember)).BlockFunc(func(s *j.Group) {
				for i := 1; i <= g.order; i++ {
					s.Case(j.Lit(i)).Block(
						j.Return(j.Qual("encoding/json", "Marshal").Call(j.Id("u").Dot(fmt.Sprintf("t%d", i)))),
					)
				}
				s.Default().Block(
					j.Panic(j.Lit("unreachable")),
				)
			})
		})

	f.Line()

	// UnmarshalJSON
	f.Func().Params(j.Id("u").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
		Id("UnmarshalJSON").
		Params(j.Id("data").Index().Byte()).Error().
		BlockFunc(func(b *j.Group) {
			for i := 1; i <= g.order; i++ {
				b.List(j.Id(fmt.Sprintf("err%d", i))).Op(":=").Qual("encoding/json", "Unmarshal").Call(
					j.Id("data"), j.Op("&").Id("u").Dot(fmt.Sprintf("t%d", i)),
				)
				b.If(j.Id(fmt.Sprintf("err%d", i)).Op("==").Nil()).Block(
					j.Id("u").Dot(g.naming.kindMember).Op("=").Lit(i),
					j.Return(j.Nil()),
				)
			}
			b.Return(j.Qual("fmt", "Errorf").Call(
				j.Lit(fmt.Sprintf("failed to unmarshal %s", g.naming.structName)),
			))
		})
}

func (g *TagUnionGenerator) fromMethodName(i int) string {
	return fmt.Sprintf("From%d", i)
}

func (g *TagUnionGenerator) GenBuilder(f *j.File) {
	// tn := g.typeNames()
	// sn := g.structName()
	// bn := g.builderName()
	// bs := g.builderStructName()

	// Builder interface
	f.Type().Id(g.naming.builderInterface).Types(g.typeNamesWithAny()...).InterfaceFunc(func(iface *j.Group) {
		for i := 1; i <= g.order; i++ {
			iface.Id(g.fromMethodName(i)).Params(
				j.Id(fmt.Sprintf("t%d", i)).Id(fmt.Sprintf("T%d", i)),
			).Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)
		}
	})

	// Builder struct
	f.Type().Id(g.naming.builderImpl).Types(g.typeNamesWithAny()...).Struct()

	// From1 ~ FromN
	for i := 1; i <= g.order; i++ {
		f.Func().Params(j.Id("b").Op("*").Id(g.naming.builderImpl).Types(g.typeNameIds()...)).
			Id(g.naming.setter(i)).
			Params(j.Id(g.naming.value(i)).Id(g.naming.typeName(i))).
			Params(j.Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
			Block(
				j.Return(j.Op("&").Id(g.naming.structName).Types(g.typeNameIds()...).Values(
					j.Dict{
						j.Id(g.naming.kindMember): j.Lit(i),
						j.Id(g.naming.value(i)):   j.Id(g.naming.value(i)),
					},
				)),
			)
	}
}

func (g *TagUnionGenerator) GenFactory(f *j.File) {
	f.Func().Id(g.naming.factoryMethod).
		Types(g.typeNamesWithAny()...).Params().
		Params(j.Op("*").Id(g.naming.builderImpl).Types(g.typeNameIds()...)).
		Block(
			j.Return(j.Op("&").Id(g.naming.builderImpl).Types(g.typeNameIds()...).Values()),
		)
}

func (g *TagUnionGenerator) Gen(f *j.File) error {
	if g.order < 2 {
		return fmt.Errorf("UnionN must have at least 2 types")
	}
	f.Comment("Code generated by goadt. DO NOT EDIT.")
	g.genStruct(f)
	g.GenWhenMethod(f)
	g.GenGetMethods(f)
	g.GenJSONMethods(f)
	g.GenBuilder(f)
	g.GenFactory(f)
	return nil
}
