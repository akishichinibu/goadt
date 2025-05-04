package tuple

import (
	j "github.com/dave/jennifer/jen"
)

const runtimePkg = "github.com/akishichinibu/goadt/pkg/runtime"

func (g *TupleGenerator) genMarshalJSON(f *j.File) {
	f.Func().
		Params(j.Id("t").Op("*").Id(g.naming.structName).Types(g.typeNameIds()...)).
		Id("MarshalJSON").
		Params().
		Params(j.Index().Byte(), j.Error()).
		Block(
			j.Id("vs").Op(":=").Index().Any().ValuesFunc(func(group *j.Group) {
				for i := 1; i <= g.order; i++ {
					group.Id("t").Dot(g.naming.value(i))
				}
			}),
			j.Return(j.Qual(runtimePkg, "MarshalJSON").Call(j.Id("vs"))),
		)
}

func (tg *TupleGenerator) genUnmarshalJSON(f *j.File) {
	f.Func().
		Params(j.Id("t").Op("*").Id(tg.naming.structName).Types(tg.typeNameIds()...)).
		Id("UnmarshalJSON").
		Params(j.Id("data").Index().Byte()).
		Params(j.Error()).
		BlockFunc(func(g *j.Group) {
			g.Var().Id("vs").Index().Any()
			g.If(j.Err().Op(":=").Qual(runtimePkg, "UnmarshalJSON").Call(j.Id("data"), j.Op("&").Id("vs")), j.Err().Op("!=").Nil()).Block(
				j.Return(j.Err()),
			)
			g.If(
				j.Len(j.Id("vs")).Op("!=").Lit(tg.order),
			).Block(
				j.Return(j.Qual("fmt", "Errorf").Call(j.Lit("expected %d values, got %d"), j.Lit(tg.order), j.Len(j.Id("vs")))),
			)

			g.Line()
			g.Var().Id("ok").Bool()
			for i := 1; i <= tg.order; i++ {
				g.List(j.Id("t").Dot(tg.naming.value(i)), j.Id("ok")).Op("=").Id("vs").Index(j.Lit(i - 1)).Assert(j.Id(tg.naming.typeName(i)))
				g.If(
					j.Op("!").Id("ok"),
				).Block(
					j.Return(j.Qual("fmt", "Errorf").Call(
						j.Lit("expected type %T, got %T"),
						j.Id("t").Dot(tg.naming.value(i)),
						j.Id("vs").Index(j.Lit(i-1)),
					)),
				)
			}
			g.Return(j.Nil())
		})
}
