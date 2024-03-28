package mapper

import (
	"reflect"
)

const tagName = "map"

// Map maps fields from struct A to struct B based on a 'map' tag.
func Map[B any](A any) B {
	var b B
	mapStruct(reflect.ValueOf(A), reflect.ValueOf(&b).Elem())
	return b
}

// mapStruct performs the actual mapping from A to B, handling nested structs.
func mapStruct(aValue reflect.Value, bValue reflect.Value) {
	aValue = reflect.Indirect(aValue)
	if aValue.Kind() != reflect.Struct {
		return
	}

	if bValue.Kind() == reflect.Ptr {
		bValue = bValue.Elem()
	}

	collector := make(map[string]reflect.Value)
	for i := 0; i < aValue.NumField(); i++ {
		field := aValue.Field(i)
		fieldType := aValue.Type().Field(i)
		tagValue, ok := fieldType.Tag.Lookup(tagName)
		if !ok {
			continue
		}
		collector[tagValue] = field
	}

	for i := 0; i < bValue.NumField(); i++ {
		bField := bValue.Field(i)
		bFieldType := bValue.Type().Field(i)
		tagValue, ok := bFieldType.Tag.Lookup(tagName)
		if !ok || !bField.CanSet() {
			continue
		}
		if aValue, exists := collector[tagValue]; exists {
			if aValue.Kind() == reflect.Struct && (bField.Kind() == reflect.Ptr || bField.Kind() == reflect.Struct) {
				// Handle nested structs: Initialize the field if it's nil.
				if bField.Kind() == reflect.Ptr && bField.IsNil() {
					bField.Set(reflect.New(bField.Type().Elem()))
				}
				mapStruct(aValue, bField)
			} else if aValue.Type().AssignableTo(bField.Type()) {
				bField.Set(aValue)
			}
		}
	}
}
