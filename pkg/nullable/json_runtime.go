package nullable

import "encoding/json"

func marshalJSON(v any) ([]byte, error) {
	if tv, ok := v.(json.Marshaler); ok {
		return tv.MarshalJSON()
	}
	return json.Marshal(v)
}

func unmarshalJSON(data []byte, v any) error {
	if tv, ok := v.(json.Unmarshaler); ok {
		return tv.UnmarshalJSON(data)
	}
	return json.Unmarshal(data, v)
}
