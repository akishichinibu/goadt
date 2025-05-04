package runtime

import "encoding/json"

func MarshalJSON(v any) ([]byte, error) {
	if tv, ok := v.(json.Marshaler); ok {
		return tv.MarshalJSON()
	}
	return json.Marshal(v)
}

func UnmarshalJSON(data []byte, v any) error {
	if tv, ok := v.(json.Unmarshaler); ok {
		return tv.UnmarshalJSON(data)
	}
	return json.Unmarshal(data, v)
}
