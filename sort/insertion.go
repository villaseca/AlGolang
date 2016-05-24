package sort

import (
	"AlGolang/utils"
	"time"
)

func InsertionSort(data []int) []int {

	defer utils.Track(time.Now(), "Insertion sorting")
	sortedArray := make([]int, 0, len(data))
	sortedArray = append(sortedArray, data[0])

	for i := 1; i < len(data); i++ {
		for j := 0; j < len(sortedArray); j++ {
			if data[i] <= sortedArray[j] {
				sortedArray = utils.Insert(sortedArray, j, data[i])
				break
			} else if (data[i] > sortedArray[j]) && (len(sortedArray)-1 == j) {
				sortedArray = append(sortedArray, data[i])
				break
			}
		}
	}

	return sortedArray
}
