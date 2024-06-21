package mapper

import (
	"github.com/shopspring/decimal"
	"testing"
	"time"
)

type StructA struct {
	FieldA              int             `map:"fieldA"`
	FieldB              string          `map:"fieldB"`
	FieldC              time.Time       `map:"fieldC"`
	FieldD              decimal.Decimal `map:"fieldD"`
	NestedStructA       NestedStructA   `map:"nestedStructA"`
	NestedSliceStructsA []NestedStructA `map:"nestedSliceStructsA"`
}

type NestedStructA struct {
	FieldA int             `map:"fieldA"`
	FieldB string          `map:"fieldB"`
	FieldC time.Time       `map:"fieldC"`
	FieldD decimal.Decimal `map:"fieldD"`
	FieldE []int           `map:"fieldE"`
}

type StructB struct {
	FieldA              int             `map:"fieldA"`
	FieldB              string          `map:"fieldB"`
	FieldC              time.Time       `map:"fieldC"`
	FieldD              decimal.Decimal `map:"fieldD"`
	NestedStructB       NestedStructB   `map:"nestedStructA"`
	NestedSliceStructsB []NestedStructB `map:"nestedSliceStructsA"`
}

type NestedStructB struct {
	FieldA int             `map:"fieldA"`
	FieldB string          `map:"fieldB"`
	FieldC time.Time       `map:"fieldC"`
	FieldD decimal.Decimal `map:"fieldD"`
	FieldE []int           `map:"fieldE"`
}

func TestMap(t *testing.T) {
	structA := []StructA{
		{
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
			NestedSliceStructsA: []NestedStructA{
				{
					FieldA: 2,
					FieldB: "hello",
					FieldC: time.Now().Add(time.Hour),
					FieldD: decimal.NewFromFloat(3.15),
					FieldE: []int{1, 2, 3},
				},
				{
					FieldA: 2,
					FieldB: "world",
					FieldC: time.Now().Add(time.Hour),
					FieldD: decimal.NewFromFloat(3.15),
					FieldE: []int{1, 2, 3},
				},
			},
		},
	}

	structB := Map[[]StructB](structA)

	if len(structA) != len(structB) {
		t.Error("invalid structB len")
	}

	for i, a := range structA {
		if a.FieldA != structB[i].FieldA {
			t.Error("invalid fieldA value")
		}

		if a.FieldB != structB[i].FieldB {
			t.Error("invalid fieldB value")
		}

		if a.FieldC != structB[i].FieldC {
			t.Error("invalid fieldC value")
		}

		if !a.FieldD.Equal(structB[i].FieldD) {
			t.Error("invalid fieldD value")
		}

		if a.NestedStructA.FieldA != structB[i].NestedStructB.FieldA {
			t.Error("invalid nestedStruct fieldA value")
		}

		if a.NestedStructA.FieldB != structB[i].NestedStructB.FieldB {
			t.Error("invalid nestedStruct fieldB value")
		}

		if a.NestedStructA.FieldC != structB[i].NestedStructB.FieldC {
			t.Error("invalid nestedStruct fieldC value")
		}

		if !a.NestedStructA.FieldD.Equal(structB[i].NestedStructB.FieldD) {
			t.Error("invalid nestedStruct fieldD value")
		}

		if len(a.NestedSliceStructsA) != len(structB[i].NestedSliceStructsB) {
			t.Error("invalid sliceNestedStruct len")
		}
	}
}
