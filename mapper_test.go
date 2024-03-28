package mapper

import "testing"

type StructA struct {
	FieldA int    `map:"fieldA"`
	FieldB string `map:"fieldB"`
}

type StructB struct {
	FieldA int    `map:"fieldA"`
	FieldB string `map:"fieldB"`
}

func TestMap(t *testing.T) {
	structA := StructA{
		FieldA: 10,
		FieldB: "hello",
	}

	structB := Map[StructB](structA)

	if structA.FieldA != structB.FieldA {
		t.Error("invalid fieldA value")
	}

	if structA.FieldB != structB.FieldB {
		t.Error("invalid fieldB value")
	}
}
