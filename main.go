package main

import (
	"AlGolang/search"
	//"AlGolang/sort"
	"AlGolang/utils"
	"fmt"
)

func main() {
	array := utils.SortedArray(100)
	num := 8
	a := search.LinearSearch(array, num)
	fmt.Println(a)
	fmt.Println(array)
}
