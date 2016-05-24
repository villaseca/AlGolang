package utils

import (
	"math/rand"
	"time"
)

func RandomArray(size int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	var array []int
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(size))
	}

	return array
}
