package aro

import (
	"reflect"
)

func Equals[T any](a, b T) bool {
	// type

	// Get type of T

	// If T is any
	// 	type of a == type of b
	// 		store type
	// 	else
	// 		return false
	// else
	// 	store type

	// Comparable, DeepEqual, Struct, and Map?

	return reflect.DeepEqual(a, b)
}
