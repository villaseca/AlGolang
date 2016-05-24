package main

import (
	"AlGolang/search"
	//"AlGolang/sort"
	"AlGolang/utils"
	"fmt"
)

func main() {
	array := utils.SortedArray(100)
	num := 88
	a := search.BinarySearch(array, num)
    if a != -1 {
        fmt.Printf("Position of %v in the array: %v\n", num, a)
    } else {
        fmt.Println("Number not found in array")
    }
}
