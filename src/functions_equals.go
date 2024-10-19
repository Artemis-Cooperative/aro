package aro

import (
	"reflect"
)

func Equals[T any](a, b T) bool {
	return reflect.DeepEqual(a, b)
}
