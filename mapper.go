package mapper

import (
	"reflect"
)

const tagName = "map"

// Map maps fields from struct A to struct B based on a 'map' tag.
// A is the source struct, and B is the target struct type.
// The function returns a new instance of B with mapped values from A.
func Map[B any](A any) B {
	aValue := reflect.Indirect(reflect.ValueOf(A))
	if aValue.Kind() != reflect.Struct {
		var zeroB B
		return zeroB // Early return with zero value of B if A is not a struct.
	}

	bContainer := new(B)
	bValue := reflect.Indirect(reflect.ValueOf(bContainer))

	collector := make(map[string]any, aValue.NumField())
	for i := 0; i < aValue.NumField(); i++ {
		field := aValue.Type().Field(i)
		tagValue, ok := field.Tag.Lookup(tagName)
		if !ok {
			continue
		}
		collector[tagValue] = aValue.Field(i).Interface()
	}

	for i := 0; i < bValue.NumField(); i++ {
		field := bValue.Type().Field(i)
		tagValue, ok := field.Tag.Lookup(tagName)
		if !ok || !bValue.Field(i).CanSet() {
			continue
		}
		if value, exists := collector[tagValue]; exists {
			newValue := reflect.ValueOf(value)
			if newValue.Type().AssignableTo(bValue.Field(i).Type()) {
				bValue.Field(i).Set(newValue)
			}
		}
	}

	return *bContainer
}
