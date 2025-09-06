package patch

import (
	"reflect"
)

func IsPatchStructEmpty(patchStruct interface{}) bool {
	val := reflect.ValueOf(patchStruct)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return true
	}

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)

		if fieldVal.CanInterface() {
			if isSetField(fieldVal.Interface()) {
				return false
			}
		}
	}
	return true
}

func isSetField(f interface{}) bool {
	val := reflect.ValueOf(f)
	if val.Kind() != reflect.Struct {
		return false
	}

	isSet := val.FieldByName("IsSet")
	if isSet.IsValid() && isSet.Kind() == reflect.Bool {
		return isSet.Bool()
	}
	return false
}
