package gostructmap

import (
	"errors"
	"reflect"
)

func MappingFrom(fromObj interface{}, toObj interface{}) error {
	if reflect.ValueOf(fromObj).Kind() != reflect.Pointer {
		return errors.New("fromObj must be a pointer")
	}
	if reflect.ValueOf(toObj).Kind() != reflect.Pointer {
		return errors.New("toObj must be a pointer")
	}
	fromType := reflect.TypeOf(fromObj)
	fromValue := reflect.ValueOf(fromObj)
	fromTypeElem := fromType.Elem()
	fromValueElem := fromValue.Elem()

	//toType := reflect.TypeOf(toObj)
	toValue := reflect.ValueOf(toObj)

	for i := 0; i < fromValueElem.NumField(); i++ {
		tagVal := fromTypeElem.Field(i).Tag.Get(TAG)
		f := reflect.Indirect(toValue).FieldByName(tagVal)
		f.SetString(fromValue.String())
	}
	return nil
}
