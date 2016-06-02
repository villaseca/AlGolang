package test

//BubbleSort takes an array and bubble sorts it
func BubbleSort(data []int) []int {

	swap := -1
	for swap != 0 {
		swap = 0
		for i := 0; i < len(data)-1; i++ {
			if data[i] > data[i+1] {
				data[i], data[i+1] = data[i+1], data[i]
				swap++
			}
		}
	}

	return data
}
