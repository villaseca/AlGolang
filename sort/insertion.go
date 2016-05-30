package sort

//InsertionSort takes an array and insertion sorts it
func InsertionSort(data []int) []int {

	sortedArray := make([]int, 0, len(data))
	sortedArray = append(sortedArray, data[0])

	for i := 1; i < len(data); i++ {
		for j := 0; j < len(sortedArray); j++ {
			if data[i] <= sortedArray[j] {
				sortedArray = insert(sortedArray, j, data[i])
				break
			} else if (data[i] > sortedArray[j]) && (len(sortedArray)-1 == j) {
				sortedArray = append(sortedArray, data[i])
				break
			}
		}
	}

	return sortedArray
}

//Insert puts value into an slice at index. Source: https://blog.golang.org/slices
func insert(slice []int, index, value int) []int {
	slice = slice[0 : len(slice)+1]
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}
