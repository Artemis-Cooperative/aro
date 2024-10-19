package aro

import (
	"testing"
)

func TestSeq(t *testing.T) {
	expected := []int{1, 2, 3, 4}
	actual := Seq(1, 4)

	if !Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%v\nActual:\t\t%v\n", expected, actual)
	}
}
