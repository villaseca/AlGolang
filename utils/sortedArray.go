package utils

import (
	"math/rand"
	"time"
)

func SortedArray(size int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	var array []int
	for i := 0; i < size; i++ {
		var noise int = rand.Intn(2)
		array = append(array, i+1+noise)
		size = size + noise
		i = i + noise
	}

	return array
}
