package nullable

const jsonNull = "null"

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	v, ok := n.Get()
	if !ok {
		return []byte(jsonNull), nil
	}
	return marshalJSON(v)
}

func (n *Nullable[T]) UnmarshalJSON(data []byte) error {
	if string(data) == jsonNull {
		n.ok = false
		return nil
	}

	var v T
	if err := unmarshalJSON(data, &v); err != nil {
		return err
	}

	n.ok = true
	n.value = v
	return nil
}
