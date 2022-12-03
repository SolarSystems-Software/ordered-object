package orderedobject

import (
	"encoding/json"
	"testing"
)

// TestMarshalWithAlphabeticKeys tests marshaling with keys being the letters of the alphabet.
//
// In encoding/json, the object used will be marshaled to `{"a":0,"b":0}`, however with our
// marshaling it should marshal into `{"b":0,"a":0}`.
func TestMarshalWithAlphabeticKeys(t *testing.T) {
	obj := new(Object[int])
	obj.Set("b", 0)
	obj.Set("a", 0)

	encoded, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	if string(encoded) != `{"b":0,"a":0}` {
		t.Fail()
	}
}

// TestMarshalWithObjects test marshaling with nested objects.
func TestMarshalWithObjects(t *testing.T) {
	tom := new(Object[any])
	tom.Set("name", "Tom")
	tom.Set("age", 21)

	alex := new(Object[any])
	alex.Set("name", "Alex")
	alex.Set("age", 25)

	people := new(Object[any])
	people.Set("tom", tom)
	people.Set("alex", alex)

	encoded, err := json.Marshal(people)
	if err != nil {
		panic(err)
	}
	if string(encoded) != `{"tom":{"name":"Tom","age":21},"alex":{"name":"Alex","age":25}}` {
		t.Fail()
	}
}
