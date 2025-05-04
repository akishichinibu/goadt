package goadt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func extraTagFromTagField(value any, tagField string) (string, error) {
	rv := reflect.ValueOf(value)
	if !rv.IsValid() {
		return "", fmt.Errorf("invalid value")
	}

	// Support pointer to struct
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return "", fmt.Errorf("cannot extract tag from nil pointer")
		}
		rv = rv.Elem()
	}

	rt := rv.Type()
	if rt.Kind() != reflect.Struct {
		return "", fmt.Errorf("value must be struct or *struct, got %s", rt.Kind())
	}

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		jsonTag := field.Tag.Get("json")
		jsonTagName := strings.Split(jsonTag, ",")[0]

		var fieldJsonName string
		if jsonTagName == "-" {
			continue // 忽略
		}
		if jsonTagName != "" {
			fieldJsonName = jsonTagName
		} else {
			// fallback to default json field name
			fieldJsonName = strings.ToLower(field.Name[:1]) + field.Name[1:]
		}

		if fieldJsonName == tagField {
			if field.Type.Kind() != reflect.String {
				return "", fmt.Errorf("field %s must be string, got %s", field.Name, field.Type.Kind())
			}
			return rv.Field(i).String(), nil
		}
	}

	return "", nil
}

type TagExtractor[T any] func(t T) string

type TagUnion2[T1 any, T2 any] struct {
	tagField string

	tag string
	t1  T1
	t2  T2

	tag1 string
	tag2 string
}

func (u *TagUnion2[T1, T2]) When(t1 func(t1 T1), t2 func(t2 T2)) {
	switch u.tag {
	case u.tagExtractor1(u.t1):
		t1(u.t1)
	case u.tagExtractor2(u.t2):
		t2(u.t2)
	default:
		panic("unreachable")
	}
}
func (u *TagUnion2[T1, T2]) As1() (T1, bool) {
	switch u.tag {
	case u.tag1:
		return u.t1, true
	default:
		return u.t1, false
	}
}
func (u *TagUnion2[T1, T2]) As2() (T2, bool) {
	if u.kind == 2 {
		return u.t2, true
	}
	return u.t2, false
}

// MarshalJSON implements the json.Marshaler interface for the Union2 type.
// It serializes the Union2 instance into JSON based on the active variant.
// If the Union2 instance holds a value of type T1, it marshals that value.
// If it holds a value of type T2, it marshals that value instead.
// An unreachable panic is triggered if the kind field has an invalid value.
func (u *TagUnion2[T1, T2]) MarshalJSON() ([]byte, error) {
	switch u.tag {
	case u.tag1:
		return json.Marshal(u.t1)
	case u.tag2:
		return json.Marshal(u.t2)
	default:
		panic("unreachable")
	}
}

func (u *TagUnion2[T1, T2]) UnmarshalJSON(data []byte) error {
	var raw map[string]any
	if err := json.Unmarshal(data, raw); err != nil {
		return err
	}
	switch raw[u.tagField] {
	case u.tag1:
		return json.Unmarshal(data, &u.t1)
	case u.tag2:
		return json.Unmarshal(data, &u.t2)
	default:
		return fmt.Errorf("failed to unmarshal Union2")
	}
}

type TagUnion2Union2Builder[T1 any, T2 any] interface {
	From1(t1 T1) *Union2[T1, T2]
	From2(t2 T2) *Union2[T1, T2]
}
type tagUnion2union2Builder[T1 any, T2 any] struct {
	tagField string
	tag1     string
	tag2     string
}

func (b *tagUnion2union2Builder[T1, T2]) From1(t1 T1) (*TagUnion2[T1, T2], error) {
	tag, err := extraTagFromTagField(t1, b.tag1)
	if err != nil {
		return nil, err
	}
	return &TagUnion2[T1, T2]{
		tagField: b.tagField,
		tag1:     b.tag1,
		tag2:     b.tag2,

		tag: tag,
		t1:  t1,
	}, nil
}

func (b *union2Builder[T1, T2]) From2(t2 T2) *TagUnion2[T1, T2] {
	return &Union2[T1, T2]{
		kind: 2,
		t2:   t2,
	}
}
func TagUnion2NewUnion2[T1 any, T2 any](
	tagField string,
	tag1 string,
	tag2 string,
) *union2Builder[T1, T2] {
	return &union2Builder[T1, T2]{}
}
