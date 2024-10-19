package aro

func Seq(min int, max int) []int {
	slice := []int{}

	for i := min; i <= max; i++ {
		slice = append(slice, i)
	}

	return slice
}
