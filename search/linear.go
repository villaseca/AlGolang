package search

func LinearSearch(data []int, num int) int {

	for i := 0; i < len(data); i++ {
		if data[i] == num {
			return i
		}
	}

	return -1
}
