package mapper

import (
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

type StructA struct {
	FieldA        int             `map:"fieldA"`
	FieldB        string          `map:"fieldB"`
	FieldC        time.Time       `map:"fieldC"`
	FieldD        decimal.Decimal `map:"fieldD"`
	NestedStructA NestedStructA   `map:"nestedStructA"`
}

type NestedStructA struct {
	FieldA int             `map:"fieldA"`
	FieldB string          `map:"fieldB"`
	FieldC time.Time       `map:"fieldC"`
	FieldD decimal.Decimal `map:"fieldD"`
}

type StructB struct {
	FieldA        int             `map:"fieldA"`
	FieldB        string          `map:"fieldB"`
	FieldC        time.Time       `map:"fieldC"`
	FieldD        decimal.Decimal `map:"fieldD"`
	NestedStructB NestedStructB   `map:"nestedStructA"`
}

type NestedStructB struct {
	FieldA int             `map:"fieldA"`
	FieldB string          `map:"fieldB"`
	FieldC time.Time       `map:"fieldC"`
	FieldD decimal.Decimal `map:"fieldD"`
}

func TestMap(t *testing.T) {
	structA := StructA{
		FieldA: 1,
		FieldB: "hello",
		FieldC: time.Now(),
		FieldD: decimal.NewFromFloat(3.14),
		NestedStructA: NestedStructA{
			FieldA: 2,
			FieldB: "world",
			FieldC: time.Now().Add(time.Hour),
			FieldD: decimal.NewFromFloat(3.15),
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

	if !structA.FieldD.Equal(structB.FieldD) {
		t.Error("invalid fieldD value")
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

	if structA.NestedStructA.FieldC != structB.NestedStructB.FieldC {
		t.Error("invalid nestedStruct fieldD value")
	}
}
