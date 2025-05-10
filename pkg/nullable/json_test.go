package nullable_test

import (
	"encoding/json"
	"testing"

	"github.com/akishichinibu/goadt/pkg/nullable"
	"github.com/stretchr/testify/require"
)

func ptr[T any](v T) *T {
	return &v
}

type jsonTestCase[T any] struct {
	name        string
	input       nullable.Nullable[T]
	exceptValue *T
	exceptJSON  string
}

func (t jsonTestCase[T]) exceptNull() bool {
	return t.exceptJSON == `null`
}

func (tc jsonTestCase[T]) run(t *testing.T) {
	data, err := json.Marshal(tc.input)
	require.NoError(t, err, tc.name)
	require.Equal(t, tc.exceptJSON, string(data), tc.name)

	var parsed nullable.Nullable[string]
	err = json.Unmarshal(data, &parsed)
	require.NoError(t, err, tc.name)

	v, ok := parsed.Get()
	switch tc.exceptNull() {
	case true:
		require.False(t, ok, tc.name)
	case false:
		require.True(t, ok, tc.name)
		require.Equal(t, *tc.exceptValue, v, tc.name)
	}
}

func TestSimpleType(t *testing.T) {
	val1 := nullable.NewWithValue("hello")
	val2 := nullable.NewNull[string]()

	cases := []jsonTestCase[string]{
		{"non-null string", val1, ptr("hello"), `"hello"`},
		{"null string", val2, nil, `null`},
	}

	for _, tc := range cases {
		tc.run(t)
	}
}

func TestSimpleTypeInStruct(t *testing.T) {
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

func TestStruct(t *testing.T) {
	type Obj struct {
		A int
		B string
	}

	val1 := nullable.NewWithValue(Obj{A: 1, B: "test"})
	val2 := nullable.NewNull[Obj]()

	cases := []jsonTestCase[Obj]{
		{"non-null struct", val1, ptr(Obj{A: 1, B: "test"}), `{"A":1,"B":"test"}`},
		{"null struct", val2, nil, `null`},
	}

	for _, tc := range cases {
		tc.run(t)
	}
}
