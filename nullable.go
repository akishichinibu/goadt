package goadt

type Nullable[T any] interface {
	Get() (T, bool)
	IsPresent() bool
}

type nullable[T any] struct {
	u *Union2[T, null]
}

func (n *nullable[T]) Get() (T, bool) {
	return n.u.As1()
}

func (n *nullable[T]) IsPresent() bool {
	_, ok := n.u.As1()
	return ok
}

func NewNull[T any]() Nullable[T] {
	v := NewUnion2[T, null]().From2(Null)
	return &nullable[T]{u: v}
}
