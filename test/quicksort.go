package test

//Based on: http://rosettacode.org/wiki/Sorting_algorithms/Quicksort

//Quicksort takes an array and quicksorts it
func Quicksort(data []int) []int {
	var less, equal, greater []int

	if len(data) > 1 {
		pivot := data[len(data)/2]
		for _, value := range data {
			if value < pivot {
				less = append(less, value)
			} else if value > pivot {
				greater = append(greater, value)
			} else {
				equal = append(equal, value)
			}
		}

		less = Quicksort(less)
		greater = Quicksort(greater)
		data = append(append(less, equal...), greater...)
	}
	return data
}
