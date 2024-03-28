package mapper

import (
	"testing"
)

type StructA struct {
	FieldA        int           `map:"fieldA"`
	FieldB        string        `map:"fieldB"`
	NestedStructA NestedStructA `map:"nestedStructA"`
}

type NestedStructA struct {
	FieldA int    `map:"fieldA"`
	FieldB string `map:"fieldB"`
}

type StructB struct {
	FieldA        int           `map:"fieldA"`
	FieldB        string        `map:"fieldB"`
	NestedStructB NestedStructB `map:"nestedStructA"`
}

type NestedStructB struct {
	FieldA int    `map:"fieldA"`
	FieldB string `map:"fieldB"`
}

func TestMap(t *testing.T) {
	structA := StructA{
		FieldA: 1,
		FieldB: "hello",
		NestedStructA: NestedStructA{
			FieldA: 2,
			FieldB: "world",
		},
	}

	structB := Map[StructB](structA)

	if structA.FieldA != structB.FieldA {
		t.Error("invalid fieldA value")
	}

	if structA.FieldB != structB.FieldB {
		t.Error("invalid fieldB value")
	}

	if structA.NestedStructA.FieldA != structB.NestedStructB.FieldA {
		t.Error("invalid nestedStruct fieldA value")
	}

	if structA.NestedStructA.FieldB != structB.NestedStructB.FieldB {
		t.Error("invalid nestedStruct fieldB value")
	}
}
