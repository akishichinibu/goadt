package tuple_test

import (
	"testing"

	"github.com/akishichinibu/goadt/pkg/tuple"
	"github.com/stretchr/testify/require"
)

func TestTuple2Getters(t *testing.T) {
	tp := tuple.NewTuple2(42, "hello")
	require.Equal(t, 42, tp.Get1())
	require.Equal(t, "hello", tp.Get2())
}

func TestTuple2Unwrap(t *testing.T) {
	tp := tuple.NewTuple2(3.14, true)
	v1, v2 := tp.Unwrap()
	require.Equal(t, 3.14, v1)
	require.Equal(t, true, v2)
}
