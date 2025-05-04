package union

import "fmt"

type unionNaming struct {
	order int

	kindMember       string
	tagMember        string
	builderInterface string
	builderImpl      string
	structName       string
	factoryMethod    string
}

func (n *unionNaming) typeName(i int) string {
	return fmt.Sprintf("T%d", i)
}

func (n *unionNaming) value(i int) string {
	return fmt.Sprintf("t%d", i)
}

func (n *unionNaming) getter(i int) string {
	return fmt.Sprintf("As%d", i)
}

func (n *unionNaming) setter(i int) string {
	return fmt.Sprintf("From%d", i)
}

func (n *unionNaming) errName(i int) string {
	return fmt.Sprintf("err%d", i)
}

func newUnionNaming(n int) unionNaming {
	return unionNaming{
		order:            n,
		kindMember:       "kind",
		builderInterface: fmt.Sprintf("Union%dBuilder", n),
		builderImpl:      fmt.Sprintf("union%dBuilder", n),
		structName:       fmt.Sprintf("Union%d", n),
		factoryMethod:    fmt.Sprintf("NewUnion%d", n),
	}
}
