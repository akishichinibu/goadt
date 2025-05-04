package nullable

import "github.com/akishichinibu/goadt/pkg/union"

type Nullable[T any] struct {
	u *union.Union2[T, null]
}

func (n Nullable[T]) Get() (T, bool) {
	return n.u.As1()
}

func (n Nullable[T]) IsPresent() bool {
	_, ok := n.u.As1()
	return ok
}

func (n Nullable[T]) IsNull() bool {
	_, ok := n.u.As2()
	return ok
}

func NewNull[T any]() Nullable[T] {
	u := union.NewUnion2[T, null]()
	v := u.From2(Null)
	return Nullable[T]{u: v}
}

func NewWithValue[T any](value T) Nullable[T] {
	u := union.NewUnion2[T, null]()
	v := u.From1(value)
	return Nullable[T]{u: v}
}
