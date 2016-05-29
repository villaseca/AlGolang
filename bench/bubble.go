//Package sort contains sorting algorithms
package bench

//BubbleSort returns a function that receives an array and bubble sorts it
func BubbleSort() func(data []int) []int {

	return func(data []int) []int {
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
}
