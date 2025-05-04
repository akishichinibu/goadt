package union_test

import (
	"encoding/json"
	"testing"

	"github.com/akishichinibu/goadt/pkg/union"
	"github.com/stretchr/testify/require"
)

func TestUnionSimpleType(t *testing.T) {
	value := union.NewUnion2[int, string]().From1(42)

	union.NewUnion2[int, string]().From1(5)

	v, ok := value.As1()
	require.True(t, ok)
	require.Equal(t, 42, v)

	_, ok = value.As2()
	require.False(t, ok)

	s, err := json.Marshal(value)
	require.NoError(t, err)
	require.Equal(t, `42`, string(s))
}

func TestUnionStructOrInterface(t *testing.T) {
	type MyStruct struct {
		Name string
	}

	value := union.NewUnion2[MyStruct, string]().From1(MyStruct{Name: "Test"})

	if v, ok := value.As1(); !ok || v.Name != "Test" {
		t.Errorf("Expected value to be MyStruct with Name 'Test', got %v", v)
	}

	if _, ok := value.As2(); ok {
		t.Errorf("Expected value to not be of type string")
	}

	s, err := json.Marshal(value)
	require.NoError(t, err)
	t.Log(string(s))
}
