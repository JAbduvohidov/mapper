package mapper

import (
	"github.com/shopspring/decimal"
	"reflect"
	"time"
)

const tagName = "map"

// Map maps fields from struct A to struct B based on a 'map' tag.
func Map[B any](A any) B {
	var b B
	aValue := reflect.ValueOf(A)
	bValue := reflect.ValueOf(&b)

	// Dereference if A is a pointer
	if aValue.Kind() == reflect.Ptr {
		if aValue.IsNil() {
			panic("source struct is nil")
		}
		aValue = aValue.Elem()
	}

	// Dereference if B is a pointer and allocate memory if needed
	if bValue.Kind() == reflect.Ptr {
		if bValue.IsNil() {
			bValue.Set(reflect.New(bValue.Type().Elem()))
		}
		bValue = bValue.Elem()
	}

	mapStruct(aValue, bValue)
	return b
}

// mapStruct performs the actual mapping from A to B, handling nested structs and slices.
func mapStruct(aValue, bValue reflect.Value) {
	if aValue.Kind() == reflect.Ptr {
		if aValue.IsNil() {
			return
		}
		aValue = aValue.Elem()
	}
	if aValue.Kind() != reflect.Struct {
		return
	}

	if bValue.Kind() == reflect.Ptr {
		if bValue.IsNil() {
			bValue.Set(reflect.New(bValue.Type().Elem()))
		}
		bValue = bValue.Elem()
	}

	collector := make(map[string]reflect.Value)
	aType := aValue.Type()
	for i := 0; i < aValue.NumField(); i++ {
		field := aValue.Field(i)
		fieldType := aType.Field(i)
		if tagValue, ok := fieldType.Tag.Lookup(tagName); ok {
			collector[tagValue] = field
		}
	}

	bType := bValue.Type()
	for i := 0; i < bValue.NumField(); i++ {
		bField := bValue.Field(i)
		bFieldType := bType.Field(i)
		if tagValue, ok := bFieldType.Tag.Lookup(tagName); ok && bField.CanSet() {
			if aField, exists := collector[tagValue]; exists {
				mapField(aField, bField)
			}
		}
	}
}

// mapField performs the actual mapping of individual fields, handling nested structs, slices, and pointers.
func mapField(aField, bField reflect.Value) {
	if !aField.IsValid() || !bField.IsValid() {
		return
	}

	if aField.Kind() == reflect.Ptr {
		if aField.IsNil() {
			return
		}
		aField = aField.Elem()
	}
	if bField.Kind() == reflect.Ptr {
		if bField.IsNil() {
			bField.Set(reflect.New(bField.Type().Elem()))
		}
		bField = bField.Elem()
	}

	switch aField.Kind() {
	case reflect.Struct:
		switch aField.Type() {
		case reflect.TypeOf(time.Time{}):
			if bField.Type() == reflect.TypeOf(time.Time{}) {
				bField.Set(aField)
			}
		case reflect.TypeOf(decimal.Decimal{}):
			if bField.Type() == reflect.TypeOf(decimal.Decimal{}) {
				bField.Set(aField)
			}
		default:
			mapStruct(aField, bField)
		}
	case reflect.Slice:
		if bField.Kind() == reflect.Slice {
			mapSlice(aField, bField)
		}
	default:
		if aField.Type().AssignableTo(bField.Type()) {
			bField.Set(aField)
		}
	}
}

// mapSlice performs the mapping of slices from A to B.
func mapSlice(aSlice, bSlice reflect.Value) {
	elemType := bSlice.Type().Elem()
	newSlice := reflect.MakeSlice(bSlice.Type(), aSlice.Len(), aSlice.Len())
	for i := 0; i < aSlice.Len(); i++ {
		aElem := aSlice.Index(i)
		bElem := reflect.New(elemType).Elem()
		mapField(aElem, bElem)
		newSlice.Index(i).Set(bElem)
	}
	bSlice.Set(newSlice)
}
