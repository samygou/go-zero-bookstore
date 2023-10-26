package tool

import "reflect"

func StructFieldCheck(origin interface{}, field string) bool {
	immutableT := reflect.TypeOf(origin)
	if _, ok := immutableT.FieldByName(field); ok {
		return true
	}

	return false
}
