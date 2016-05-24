package main

import (
	"AlGolang/search"
	"AlGolang/sort"
	"AlGolang/utils"
	"fmt"
)

func main() {
	array := utils.RandomArray(10000)
	num := 542
	sort.InsertionSort(array)
	sort.BubbleSort(array)
	sort.SelectionSort(array)

	d := search.LinearSearch(array, num)

	if d != -1 {
		fmt.Printf("Position of %v in the array: %v\n", num, d)
	} else {
		fmt.Println("Number not found in array")
	}

}
