package nullable_test

import (
	"encoding/json"
	"testing"

	"github.com/akishichinibu/goadt/pkg/nullable"
	"github.com/stretchr/testify/require"
)

type testCase1[T any] struct {
	name     string
	input    nullable.Nullable[T]
	expected string
}

func (t testCase1[T]) ExceptedNull() bool {
	return t.expected == `null`
}

func TestSingleValue(t *testing.T) {

	val1 := nullable.NewWithValue("hello")
	val2 := nullable.NewNull[string]()

	cases := []testCase1[string]{
		{"non-null string", val1, `"hello"`},
		{"null string", val2, `null`},
	}

	for _, c := range cases {
		data, err := json.Marshal(c.input)
		require.NoError(t, err, c.name)
		require.JSONEq(t, c.expected, string(data), c.name)

		var parsed nullable.Nullable[string]
		err = json.Unmarshal(data, &parsed)
		require.NoError(t, err, c.name)

		v, ok := parsed.Get()
		switch c.ExceptedNull() {
		case true:
			require.False(t, ok, c.name)
			require.Equal(t, "", v, c.name)
		case false:
			require.True(t, ok, c.name)
			require.Equal(t, "hello", v, c.name)
		}
	}
}

func TestNullableAsField(t *testing.T) {
	type Obj struct {
		Key nullable.Nullable[string] `json:"key"`
	}

	obj1 := Obj{Key: nullable.NewWithValue("value")}
	obj2 := Obj{Key: nullable.NewNull[string]()}

	data1, err := json.Marshal(obj1)
	require.NoError(t, err)
	require.JSONEq(t, `{"key":"value"}`, string(data1))

	data2, err := json.Marshal(obj2)
	require.NoError(t, err)
	require.JSONEq(t, `{"key":null}`, string(data2))

	var parsed1, parsed2 Obj
	require.NoError(t, json.Unmarshal(data1, &parsed1))
	require.NoError(t, json.Unmarshal(data2, &parsed2))

	v1, ok1 := parsed1.Key.Get()
	require.True(t, ok1)
	require.Equal(t, "value", v1)

	_, ok2 := parsed2.Key.Get()
	require.False(t, ok2)
}
