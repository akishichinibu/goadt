package union_test

import (
	"encoding/json"
	"testing"

	"github.com/akishichinibu/goadt/pkg/union"
	"github.com/stretchr/testify/require"
)

func TestMarshalUnion(t *testing.T) {
	type MyStruct struct {
		Name string
	}

	value := union.NewUnion2[MyStruct, string]().From1(MyStruct{Name: "Test"})

	data, err := json.Marshal(value)
	require.NoError(t, err)

	var unmarshaled union.Union2[MyStruct, string]
	err = json.Unmarshal(data, &unmarshaled)
	require.NoError(t, err)

	v, ok := unmarshaled.As1()
	require.True(t, ok)
	require.Equal(t, "Test", v.Name)
}
