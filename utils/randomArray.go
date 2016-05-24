package utils

import (
        "math/rand"
        "time"
)

func RandomArray(size int) []int {
    rand.Seed(time.Now().UTC().UnixNano())
	array := make([]int, size)
    for i := 0; i < size; i++ {
        array[i] = rand.Intn(size)        
    }
    
    return array
}