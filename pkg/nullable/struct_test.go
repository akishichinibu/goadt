package nullable_test

import "fmt"

type testInterface interface {
	MyInterface()
	GetA() int
	GetB() string

	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

type testStruct struct {
	a int
	b string
}

var _ testInterface = &testStruct{}

func (t *testStruct) MyInterface() {}

func (t *testStruct) GetA() int {
	return t.a
}

func (t *testStruct) GetB() string {
	return t.b
}

func (t *testStruct) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf(`{"a":%d,"b":"%s"}`, t.a, t.b)
	return []byte(s), nil
}

func (t *testStruct) UnmarshalJSON(data []byte) error {
	// Simulate unmarshalling
	t.a = 1
	t.b = "test"
	return nil
}
