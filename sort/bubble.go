package sort

import (
	"AlGolang/utils"
	"time"
)

func BubbleSort(data []int) []int {

	defer utils.Track(time.Now(), "Bubble sorting")
	sortedArray := make([]int, len(data))
	copy(sortedArray, data)

	swap := -1
	aux := 0
	for swap != 0 {
		swap = 0
		for i := 0; i < len(sortedArray)-1; i++ {
			if sortedArray[i] > sortedArray[i+1] {
				aux = sortedArray[i]
				sortedArray[i] = sortedArray[i+1]
				sortedArray[i+1] = aux
				swap++
			}
		}
	}

	return sortedArray
}
