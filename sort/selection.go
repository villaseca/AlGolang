package sort

import (
	"AlGolang/utils"
	"time"
)

func SelectionSort(data []int) []int {

	defer utils.Track(time.Now(), "Selection sorting")
	sortedArray := make([]int, len(data))
	copy(sortedArray, data)

	loc := 0
	min := 0

	for i := 0; i < len(sortedArray); i++ {
		min = sortedArray[i]
		for j := i; j < len(sortedArray); j++ {
			if sortedArray[j] <= min {
				min = sortedArray[j]
				loc = j
			}
		}
		sortedArray[loc] = sortedArray[i]
		sortedArray[i] = min
	}
	return sortedArray
}
