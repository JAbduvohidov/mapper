package mapper

import (
	"testing"
	"time"
)

type StructA struct {
	FieldA        int           `map:"fieldA"`
	FieldB        string        `map:"fieldB"`
	FieldC        time.Time     `map:"fieldC"`
	NestedStructA NestedStructA `map:"nestedStructA"`
}

type NestedStructA struct {
	FieldA int       `map:"fieldA"`
	FieldB string    `map:"fieldB"`
	FieldC time.Time `map:"fieldC"`
}

type StructB struct {
	FieldA        int           `map:"fieldA"`
	FieldB        string        `map:"fieldB"`
	FieldC        time.Time     `map:"fieldC"`
	NestedStructB NestedStructB `map:"nestedStructA"`
}

type NestedStructB struct {
	FieldA int       `map:"fieldA"`
	FieldB string    `map:"fieldB"`
	FieldC time.Time `map:"fieldC"`
}

func TestMap(t *testing.T) {
	structA := StructA{
		FieldA: 1,
		FieldB: "hello",
		FieldC: time.Now(),
		NestedStructA: NestedStructA{
			FieldA: 2,
			FieldB: "world",
			FieldC: time.Now().Add(time.Hour),
		},
	}

	structB := Map[StructB](structA)

	if structA.FieldA != structB.FieldA {
		t.Error("invalid fieldA value")
	}

	if structA.FieldB != structB.FieldB {
		t.Error("invalid fieldB value")
	}

	if structA.FieldC != structB.FieldC {
		t.Error("invalid fieldC value")
	}

	if structA.NestedStructA.FieldA != structB.NestedStructB.FieldA {
		t.Error("invalid nestedStruct fieldA value")
	}

	if structA.NestedStructA.FieldB != structB.NestedStructB.FieldB {
		t.Error("invalid nestedStruct fieldB value")
	}

	if structA.NestedStructA.FieldC != structB.NestedStructB.FieldC {
		t.Error("invalid nestedStruct fieldC value")
	}
}
