package main

import (
	"AlGolang/search"
	"AlGolang/sort"
	"AlGolang/utils"
	"fmt"
)

func main() {
	array := utils.RandomArray(50)
	num := 33
	fmt.Println(array)
	b := sort.SelectionSort(array)
	fmt.Println(b)
	a := search.BinarySearch(array, num)

	if a != -1 {
		fmt.Printf("Position of %v in the array: %v\n", num, a)
	} else {
		fmt.Println("Number not found in array")
	}

}
