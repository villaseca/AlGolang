package test

//Based on: http://rosettacode.org/wiki/Sorting_algorithms/Insertion_sort

//InsertionSort takes an array and insertion sorts it
func InsertionSort(data []int) []int {

	for i := 1; i < len(data); i++ { //first is considered sorted
		value := data[i]                    //next element to sort
		j := i - 1                          //last index in sorted subarray
		for (j >= 0) && (data[j] > value) { //if unsorted element is smaller than the last sorted element
			data[j+1] = data[j] //move that sorted element to the right
			j = j - 1           //move to previous index in sorted subarray
		}
		data[j+1] = value //insert element into sorted subarray
	}

	return data //return
}
