package search

func BinarySearch(data []int, num int) int {

	start := 0
	end := len(data) - 1

	for end >= start {

		middle := (start + end) / 2

		if data[middle] == num {
			return middle
		} else if num < data[middle] {
			end = middle - 1
		} else if num > data[middle] {
			start = middle + 1
		}
	}
	return -1

}
