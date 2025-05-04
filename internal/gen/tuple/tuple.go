package tuple

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
)

type TupleGenerator struct {
	order  int
	naming tupleNaming
}

func NewTupleGenerator(order int) *TupleGenerator {
	return &TupleGenerator{
		order:  order,
		naming: newTupleNaming(order),
	}
}

func (g *TupleGenerator) typeNameIds() (cs []j.Code) {
	for i := range g.order {
		cs = append(cs, j.Id(g.naming.typeName(i+1)))
	}
	return cs
}

func (g *TupleGenerator) paramsWithTypes() (cs []j.Code) {
	for i := 1; i <= g.order; i++ {
		cs = append(cs, j.Id(g.naming.value(i)).Id(g.naming.typeName(i)))
	}
	return cs
}

func (g *TupleGenerator) typeNamesWithAny() (cs []j.Code) {
	for id := range g.typeNameIds() {
		cs = append(cs, j.Id(g.naming.typeName(id+1)).Any())
	}
	return cs
}

func (g *TupleGenerator) genStruct(f *j.File) {
	f.Commentf("%s is a generic tuple type that holds %d values of types Tn.", g.naming.structName, g.order)
	f.Type().Id(g.naming.structName).Types(g.typeNamesWithAny()...).StructFunc(func(s *j.Group) {
		for i := 1; i <= g.order; i++ {
			s.Id(g.naming.value(i)).Id(g.naming.typeName(i))
		}
	})
}

func (g *TupleGenerator) GenGetters(f *j.File) {
	for i := 1; i <= g.order; i++ {
		f.Commentf("Get%d returns the %dth value in the tuple.", i, i)
		f.Func().Params(j.Id("t").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
			Id(g.naming.getter(i)).Params().
			Params(j.Id(g.naming.typeName(i))).
			Block(
				j.Return(j.Id("t").Dot(g.naming.value(i))),
			)
	}
}

func (g *TupleGenerator) genUnwrap(f *j.File) {
	f.Commentf("Unwrap returns all values in the tuple.")
	f.Func().Params(j.Id("t").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
		Id("Unwrap").Params().
		ParamsFunc(func(p *j.Group) {
			for i := 1; i <= g.order; i++ {
				p.Id(g.naming.typeName(i))
			}
		}).
		BlockFunc(func(b *j.Group) {
			returns := make([]j.Code, 0, g.order)
			for i := 1; i <= g.order; i++ {
				returns = append(returns, j.Id("t").Dot(g.naming.value(i)))
			}
			b.Return(returns...)
		})
}

func (g *TupleGenerator) genFactory(f *j.File) {
	f.Func().Id(g.naming.builderName).
		Types(g.typeNamesWithAny()...).
		Params(g.paramsWithTypes()...).
		Params(j.Id(g.naming.structName).Types(g.typeNameIds()...)).
		Block(
			j.Return(j.Id(g.naming.structName).Types(g.typeNameIds()...).ValuesFunc(func(v *j.Group) {
				for i := 1; i <= g.order; i++ {
					v.Id(g.naming.value(i)).Op(":").Id(g.naming.value(i))
				}
			})),
		)
}

func (g *TupleGenerator) Gen(f *j.File) error {
	if g.order < 2 {
		return fmt.Errorf("TupleN must have at least 2 elements")
	}
	g.genStruct(f)
	g.GenGetters(f)
	g.genUnwrap(f)
	g.genFactory(f)
	g.genMarshalJSON(f)
	g.genUnmarshalJSON(f)
	return nil
}
