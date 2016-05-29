//Package utils has utility small functions to make the algorithms more readable
package utils

import "math/rand"

//RandomArray generates an unordered array of integers
func RandomArray(size int, seed int64) []int {
	rand.Seed(seed)
	var array []int
	for i := 0; i < size; i++ {
		array = append(array, rand.Intn(size))
	}

	return array
}
