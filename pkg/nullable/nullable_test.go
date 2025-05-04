package nullable_test

import (
	"testing"

	"github.com/akishichinibu/goadt/pkg/nullable"
	"github.com/stretchr/testify/require"
)

func TestNewNull(t *testing.T) {
	n := nullable.NewNull[int]()
	val, ok := n.Get()

	require.False(t, ok, "expected null to not contain value")
	require.Zero(t, val, "expected default zero value")
	require.False(t, n.IsPresent(), "expected IsPresent to be false")
}

func TestNewNullable(t *testing.T) {
	n := nullable.NewWithValue(42)
	val, ok := n.Get()

	require.True(t, ok, "expected nullable to contain value")
	require.Equal(t, 42, val)
	require.True(t, n.IsPresent(), "expected IsPresent to be true")
}

func TestNullEquality(t *testing.T) {
	n1 := nullable.NewNull[int]()
	n2 := nullable.NewNull[int]()

	require.Equal(t, n1, n2, "expected two nulls to be equal")
}
