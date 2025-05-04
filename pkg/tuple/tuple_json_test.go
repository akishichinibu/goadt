package tuple_test

import (
	"encoding/json"
	"testing"

	"github.com/akishichinibu/goadt/pkg/tuple"
	"github.com/stretchr/testify/require"
)

func TestTuple2JSON(t *testing.T) {
	original := tuple.NewTuple2(123.4, "abc")

	data, err := json.Marshal(&original)
	require.NoError(t, err)
	t.Log(string(data))

	expected := `[123.4,"abc"]`
	require.Equal(t, expected, string(data), "expected JSON string does not match")

	var decoded tuple.Tuple2[float64, string]
	require.NoError(t, json.Unmarshal(data, &decoded))

	v1, v2 := decoded.Unwrap()
	require.Equal(t, original.Get1(), v1)
	require.Equal(t, original.Get2(), v2)
}
