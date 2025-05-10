package nullable

type Nullable[T any] struct {
	ok    bool
	value T
}

func (n Nullable[T]) Get() (T, bool) {
	if n.IsPresent() {
		return n.value, true
	}
	return n.value, false
}

func (n Nullable[T]) IsPresent() bool {
	return n.ok
}

func (n Nullable[T]) IsNull() bool {
	return !n.ok
}

func NewNull[T any]() Nullable[T] {
	return Nullable[T]{ok: false}
}

func NewWithValue[T any](value T) Nullable[T] {
	return Nullable[T]{ok: true, value: value}
}
