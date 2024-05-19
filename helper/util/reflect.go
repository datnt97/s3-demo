package util

import "reflect"

func IsZeroOfUnderlyingType(v interface{}) bool {
	return v == nil || reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
