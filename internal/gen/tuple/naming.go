package tuple

import "fmt"

type tupleNaming struct {
	order int

	structName  string
	builderName string
}

func (n *tupleNaming) typeName(i int) string {
	return fmt.Sprintf("T%d", i)
}

func (n *tupleNaming) value(i int) string {
	return fmt.Sprintf("t%d", i)
}

func (n *tupleNaming) getter(i int) string {
	return fmt.Sprintf("Get%d", i)
}

func newTupleNaming(n int) tupleNaming {
	return tupleNaming{
		order: n,

		structName:  fmt.Sprintf("Tuple%d", n),
		builderName: fmt.Sprintf("NewTuple%d", n),
	}
}
