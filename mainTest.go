package main

import (
	"AlGolang/test"
	"AlGolang/utils"
	"fmt"
)

func main() {

	array := utils.RandomArray(100, 37)
	fmt.Println("Unsorted", array)

	fmt.Println("Sorted", test.Quicksort(array))

}
