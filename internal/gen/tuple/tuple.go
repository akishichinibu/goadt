package tuple

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
)

type TupleGenerator struct {
	n int

	naming tupleNaming
}

func NewTupleGenerator(n int) *TupleGenerator {
	return &TupleGenerator{
		n: n,

		naming: newTupleNaming(n),
	}
}

func (g *TupleGenerator) typeNameIds() (cs []j.Code) {
	for i := range g.n {
		cs = append(cs, j.Id(g.naming.typeName(i+1)))
	}
	return cs
}

func (g *TupleGenerator) paramsWithTypes() (cs []j.Code) {
	for i := 1; i <= g.n; i++ {
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

func (g *TupleGenerator) GenStruct(f *j.File) {
	f.Type().Id(g.naming.structName).Types(g.typeNamesWithAny()...).StructFunc(func(s *j.Group) {
		for i := 1; i <= g.n; i++ {
			s.Id(g.naming.value(i)).Id(g.naming.typeName(i))
		}
	})
}

func (g *TupleGenerator) GenMethods(f *j.File) {
	for i := 1; i <= g.n; i++ {
		f.Func().Params(j.Id("t").Op("*").Id(g.naming.structName)).Add(g.typeNameIds()...).
			Id(g.naming.getter(i)).Params().
			Params(j.Id(g.naming.typeName(i))).
			Block(
				j.Return(j.Id("t").Dot(g.naming.value(i))),
			)
	}

	f.Func().Params(j.Id("t").Op("*").Id(g.naming.structName)).Add(g.typeNameIds()...).
		Id("Unwrap").Params().
		ParamsFunc(func(p *j.Group) {
			for i := 1; i <= g.n; i++ {
				p.Id(g.naming.typeName(i))
			}
		}).
		BlockFunc(func(b *j.Group) {
			returns := make([]j.Code, 0, g.n)
			for i := 1; i <= g.n; i++ {
				returns = append(returns, j.Id("t").Dot(g.naming.value(i)))
			}
			b.Return(returns...)
		})
}

func (g *TupleGenerator) GenFactory(f *j.File) {
	f.Func().Id(g.naming.builderName).
		Types(g.typeNameIds()...).
		Params(g.paramsWithTypes()...).
		Params(j.Id(g.naming.structName).Types(g.typeNameIds()...)).
		Block(
			j.Return(j.Id(g.naming.structName).Types(g.typeNameIds()...).ValuesFunc(func(v *j.Group) {
				for i := 1; i <= g.n; i++ {
					v.Id(g.naming.value(i)).Op(":").Id(g.naming.value(i))
				}
			})),
		)
}

func (g *TupleGenerator) Gen(f *j.File) error {
	if g.n < 2 {
		return fmt.Errorf("TupleN must have at least 2 elements")
	}
	g.GenStruct(f)
	g.GenMethods(f)
	g.GenFactory(f)
	return nil
}
