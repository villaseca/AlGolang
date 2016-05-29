package utils

import "math/rand"

//SortedArray generates an ordered array of integers
func SortedArray(size int, seed int64) []int {
	rand.Seed(seed)
	var array []int
	for i := 0; i < size; i++ {
		noise := rand.Intn(2)
		array = append(array, i+1+noise)
		size = size + noise
		i = i + noise
	}

	return array
}
