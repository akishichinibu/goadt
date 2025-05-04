package nullable

import (
	"encoding/json"

	"github.com/akishichinibu/goadt/pkg/union"
)

const jsonNull = "null"

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	v, ok := n.Get()
	if !ok {
		return []byte(jsonNull), nil
	}
	return json.Marshal(v)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	u := union.NewUnion2[T, null]()

	if string(data) == jsonNull {
		n.u = u.From2(Null)
		return nil
	}

	var v T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	n.u = u.From1(v)
	return nil
}
