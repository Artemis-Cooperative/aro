package aro

import(
	"reflect"
)

// Concatenate the slices given
func Concat[T any](slices ...[]T) []T {
	var concatenated []T

	for _, slice := range slices {
		concatenated = append(concatenated, slice...)
	}

	return concatenated
}

// Panic if the argument given is not a slice
func ExpectSlice(v any) {
	ExpectKind(v, "slice")
}

// Return a slice of elements that match the criteria given.
func Filter[F any](filterables []F, criteria func(F) bool) []F {
	var filtered []F

	for _, filterable := range filterables {
		if criteria(filterable) {
			filtered = append(filtered, filterable)
		}
	}

	return filtered
}

// Return the type of the elements of the slice given
func ItemType(slice any) reflect.Kind {
	ExpectSlice(slice)

	return reflect.
		TypeOf(slice).
		Elem().
		Kind()
}

// Return a sequential list of indexes based on the length of the slice given
func Indexes[T any](slice []T) []int {
	return Seq(0, len(slice)-1)
}

// Return true if the data given is a slice. Else, false.
func IsSlice(slice any) bool {
	return HasKind(slice, "slice")
}

// Prepend the item given to the slice given
func Prepend[T any](slice []T, item T) []T {
	return append([]T{item}, slice...)
}

// Remove an element from the slice by its index
func Remove[T any](slice []T, i int) []T {
	return append(slice[:i], slice[(i+1):]...)
}

// Remove all empty slices from the 2D slice given
func RemoveEmpty[T any](slices [][]T) [][]T {
	emptyIndexes := []int{}

	// Determine which slices are empty
	for i := range slices {
		if len(slices[i]) == 0 {
			emptyIndexes = Prepend(emptyIndexes, i)
		}
	}

	// Remove those slices
	for _, index := range emptyIndexes {
		slices = Remove(slices, index)
	}

	return slices
}

// Reverse the slice given
func Reverse[T any](slice []T) []T {
	reversed := []T{}

	for i := (len(slice) - 1); i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}

// Remove the first element from the slice given
func Shift[T any](slice []T) []T {
	switch len(slice) {
	case 0:
		return slice
	case 1:
		return []T{}
	default:
		return slice[1:]
	}
}