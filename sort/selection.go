package sort

func SelectionSort(data []int) []int {
    
    loc := 0
    min := 0
    
    for i := 0; i < len(data); i++ {
        min = data[i]
        for j := i; j < len(data); j++ {
            if data[j] <= min {
                min = data[j]
                loc = j
            }
        }
        data[loc] = data[i]
        data[i] = min
    }
	return data
}
