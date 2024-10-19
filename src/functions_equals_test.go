package aro

import (
	"strings"
	"testing"
)

func TestEquals(t *testing.T) {
	testCases := []struct {
		name     string
		args     []any
		expected bool
	}{
		{
			name:     "empty strings are equal",
			args:     []any{"", ""},
			expected: true,
		},
		{
			name:     "space vs empty string is not equal",
			args:     []any{" ", ""},
			expected: false,
		},
		{
			name:     "unequal strings are not equal",
			args:     []any{"a", "b"},
			expected: false,
		},
		{
			name:     "equal characters in unequal order are not equal",
			args:     []any{"ab", "ba"},
			expected: false,
		},
		{
			name:     "equal strings with unequal spaces are not equal",
			args:     []any{"a ", "a"},
			expected: false,
		},
		{
			name:     "unequal trimable whitespaces are not equal",
			args:     []any{"a ", " a"},
			expected: false,
		},
		{
			name:     "unequal untrimmable whitespaces are not equal",
			args:     []any{"a a", "a  a"},
			expected: false,
		},
		{
			name:     "equal booleans are equal",
			args:     []any{true, true},
			expected: true,
		},
		{
			name:     "unequal booleans are not equal",
			args:     []any{true, false},
			expected: false,
		},
		{
			name:     "equal integers are equal",
			args:     []any{1, 1},
			expected: true,
		},
		{
			name:     "unequal integers are not equal",
			args:     []any{1, 0},
			expected: false,
		},
		{
			name:     "unequal signs are not equal",
			args:     []any{1, -1},
			expected: false,
		},
		{
			name:     "equal floats are equal",
			args:     []any{1.0, 1.0},
			expected: true,
		},
		{
			name:     "equal float vs int is not equal",
			args:     []any{float32(1.0), int(1)},
			expected: false,
		},
		{
			name:     "unequal floats are not equal",
			args:     []any{1.0, 1.1},
			expected: false,
		},
		{
			name:     "empty slices are equal",
			args:     []any{[]int{}, []int{}},
			expected: true,
		},
		{
			name:     "not empty slice vs empty slice is not equal",
			args:     []any{[]int{1}, []int{}},
			expected: false,
		},
		{
			name:     "equal slices are equal",
			args:     []any{[]int{1, 2}, []int{1, 2}},
			expected: true,
		},
		{
			name:     "unequal value orders are not equal",
			args:     []any{[]int{1, 2}, []int{2, 1}},
			expected: false,
		},
		{
			name:     "unequal slice args are not equal",
			args:     []any{[]int{1, 2}, []int{1, 3}},
			expected: false,
		},
		{
			name:     "unequal slices lengths are not equal",
			args:     []any{[]int{1, 2}, []int{1, 2, 3}},
			expected: false,
		},
		{
			name:     "empty slices of unequal types are not equal",
			args:     []any{[]int{}, []float32{}},
			expected: false,
		},
		{
			name:     "empty maps are equal",
			args:     []any{map[int]int{}, map[int]int{}},
			expected: true,
		},
		{
			name:     "unequal typed keys are not equal",
			args:     []any{map[int]int{}, map[string]int{}},
			expected: false,
		},
		{
			name:     "unequal typed args are not equal",
			args:     []any{map[int]int{}, map[int]string{}},
			expected: false,
		},
		{
			name:     "unequal maps are not equal",
			args:     []any{map[int]int{1: 1}, map[int]int{}},
			expected: false,
		},
		{
			name:     "equal maps are equal",
			args:     []any{map[int]int{1: 1}, map[int]int{1: 1}},
			expected: true,
		},
		{
			name:     "unequal map keys are not equal",
			args:     []any{map[int]int{1: 1}, map[int]int{2: 1}},
			expected: false,
		},
		{
			name:     "unequal map args are not equal",
			args:     []any{map[int]int{1: 1}, map[int]int{1: 2}},
			expected: false,
		},
		{
			name:     "unequal lengths are not equal",
			args:     []any{map[int]int{1: 1}, map[int]int{1: 1, 2: 1}},
			expected: false,
		},
		{
			name: "equal structs are equal",
			args: []any{
				struct {
					field1 int
					field2 string
				}{
					field1: 1,
					field2: "2",
				},
				struct {
					field1 int
					field2 string
				}{
					field1: 1,
					field2: "2",
				}},
			expected: true,
		},
		// TODO: Contribute to reflect package to change
		// {
		// 	name: "unequal field orders are equal",
		// 	args: []any{
		// 		struct {
		// 			field1 int
		// 			field2 string
		// 		}{
		// 			field1: 1,
		// 			field2: "2",
		// 		},
		// 		struct {
		// 			field2 string
		// 			field1 int
		// 		}{
		// 			field1: 1,
		// 			field2: "2",
		// 		}},
		// 	expected: true,
		// },
		{
			name: "unequal instantiation orders are equal",
			args: []any{
				struct {
					field1 int
					field2 string
				}{
					field1: 1,
					field2: "2",
				},
				struct {
					field1 int
					field2 string
				}{
					field2: "2",
					field1: 1,
				}},
			expected: true,
		},
		{
			name: "unequal struct types are not equal",
			args: []any{
				struct {
					field1 int8
					field2 int16
				}{
					field1: 1,
					field2: 2,
				},
				struct {
					field1 int16
					field2 int8
				}{
					field1: 1,
					field2: 2,
				}},
			expected: false,
		},
		{
			name: "unequal struct keys are not equal",
			args: []any{
				struct {
					field1 int
					field2 string
				}{
					field1: 1,
					field2: "2",
				},
				struct {
					field3 int
					field4 string
				}{
					field3: 1,
					field4: "2",
				}},
			expected: false,
		},
		{
			name: "unequal struct args are not equal",
			args: []any{
				struct {
					field1 int
					field2 int
				}{
					field1: 1,
					field2: 2,
				},
				struct {
					field1 int
					field2 int
				}{
					field1: 2,
					field2: 1,
				}},
			expected: false,
		},
		{
			name: "unequal struct lengths are not equal",
			args: []any{
				struct {
					field1 int
					field2 string
				}{
					field1: 1,
					field2: "2",
				},
				struct {
					field1 int
					field2 string
					field3 string
				}{
					field1: 1,
					field2: "2",
					field3: "3",
				}},
			expected: false,
		},
	}
	failedTestCases := []string{}

	for _, testCase := range testCases {
		if testCase.expected != Equals(testCase.args[0], testCase.args[1]) {
			failedTestCases = append(failedTestCases, testCase.name)
		}
	}

	if len(failedTestCases) > 0 {
		t.Errorf("Failed test cases: \n" + pretty(failedTestCases))
	}
}

func pretty(a []string) string {
	out := "[\n    "

	out += strings.Join(a, "\n    ")

	out += "\n]\n"

	return out
}
