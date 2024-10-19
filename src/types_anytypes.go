package aro

import (
	"reflect"
)

func ExpectKind(v any, expected string) {
	actual := KindString(v)

	if actual != expected {
		panic("expected a " + expected + ", but received a(n) " + actual)
	}
}

func KindString(v any) string {
	return reflect.TypeOf(v).Kind().String()
}

// Return true if the kind of the value given equals the expected kind given. Else, false.
func HasKind(v any, expected string) bool {
	return KindString(v) == expected
}
