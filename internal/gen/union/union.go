package union

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
)

const runtimePkg = "github.com/akishichinibu/goadt/pkg/union"

type UnionGenerator struct {
	order  int
	naming unionNaming
}

func NewUnionGenerator(n int) *UnionGenerator {
	return &UnionGenerator{
		order:  n,
		naming: newUnionNaming(n),
	}
}

func (g *UnionGenerator) typeNameIds() (cs []j.Code) {
	for i := range g.order {
		cs = append(cs, j.Id(g.naming.typeName(i+1)))
	}
	return cs
}

func (g *UnionGenerator) typeNamesWithAny() (cs []j.Code) {
	for id := range g.typeNameIds() {
		cs = append(cs, j.Id(g.naming.typeName(id+1)).Any())
	}
	return cs
}

func (g *UnionGenerator) paramsWithSingleParamFunc() (cs []j.Code) {
	for i := 1; i <= g.order; i++ {
		cs = append(cs, j.Id(g.naming.value(i)).Func().Params(
			j.Id(g.naming.value(i)).Id(g.naming.typeName(i)),
		))
	}
	return cs
}

func (g *UnionGenerator) genStruct(f *j.File) {
	f.Commentf("Union%d is a generic tagged union type that holds a value of the type Tn.", g.order)
	f.Type().Id(g.naming.structName).Types(g.typeNamesWithAny()...).StructFunc(func(s *j.Group) {
		s.Id(g.naming.kindMember).Uint8()
		for i := 1; i <= g.order; i++ {
			s.Id(g.naming.value(i)).Id(g.naming.typeName(i))
		}
	})
}

func (g *UnionGenerator) genWhenMethod(f *j.File) {
	f.Commentf("When invokes one of the provided functions depending on the stored variant.")
	f.Commentf("If the value is of type Tn, calls tn.")

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
func (g *UnionGenerator) genGetMethods(f *j.File) {
	for i := 1; i <= g.order; i++ {
		f.Commentf("As%d returns the value as T%d and a boolean indicating whether the value is indeed T%d.", i, i, i)
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

func (g *UnionGenerator) genJSONMethods(f *j.File) {
	f.Commentf("MarshalJSON implements the `json.Marshaler` interface for %s.", g.naming.structName)
	f.Func().Params(j.Id("u").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
		Id("MarshalJSON").
		Params().Params(j.Index().Byte(), j.Error()).
		BlockFunc(func(b *j.Group) {
			b.Switch(j.Id("u").Dot(g.naming.kindMember)).BlockFunc(func(s *j.Group) {
				for i := 1; i <= g.order; i++ {
					s.Case(j.Lit(i)).Block(
						j.Return(j.Id("MarshalJSON").Call(j.Id("u").Dot(g.naming.value(i)))),
					)
				}
				s.Default().Block(
					j.Panic(j.Lit("unreachable")),
				)
			})
		})

	f.Line()

	// UnmarshalJSON
	f.Commentf("UnmarshalJSON implements the `json.Unmarshaler` interface for %s.", g.naming.structName)
	f.Commentf("It attempts to decode the JSON data into T1 or Tn.")
	f.Func().Params(j.Id("u").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
		Id("UnmarshalJSON").
		Params(j.Id("data").Index().Byte()).Error().
		BlockFunc(func(b *j.Group) {
			for i := 1; i <= g.order; i++ {
				b.List(j.Id(g.naming.errName(i))).Op(":=").Id("UnmarshalJSON").Call(
					j.Id("data"), j.Op("&").Id("u").Dot(g.naming.value(i)),
				)
				b.If(j.Id(g.naming.errName(i)).Op("==").Nil()).Block(
					j.Id("u").Dot(g.naming.kindMember).Op("=").Lit(i),
					j.Return(j.Nil()),
				)
			}
			b.Return(j.Qual("fmt", "Errorf").Call(
				j.Lit(fmt.Sprintf("failed to unmarshal %s", g.naming.structName)),
			))
		})
}

func (g *UnionGenerator) genBuilder(f *j.File) {
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

func (g *UnionGenerator) genFactory(f *j.File) {
	f.Commentf("NewUnion%d returns a builder to construct %s values using From1 or From%d.", g.order, g.naming.structName, g.order)
	f.Func().Id(g.naming.factoryMethod).
		Types(g.typeNamesWithAny()...).Params().
		Params(j.Op("*").Id(g.naming.builderImpl).Types(g.typeNameIds()...)).
		Block(
			j.Return(j.Op("&").Id(g.naming.builderImpl).Types(g.typeNameIds()...).Values()),
		)
}

func (g *UnionGenerator) Gen(f *j.File) error {
	if g.order < 2 {
		return fmt.Errorf("UnionN must have at least 2 types")
	}

	g.genStruct(f)
	g.genWhenMethod(f)
	g.genGetMethods(f)
	g.genJSONMethods(f)
	g.genBuilder(f)
	g.genFactory(f)
	return nil
}
