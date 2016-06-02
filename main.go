package main

import (
	"AlGolang/bench"
	"AlGolang/utils"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

var arraySize int
var testSize int

func main() {
	fmt.Printf("Array size: ")
	fmt.Scanf("%d", &arraySize)
	fmt.Printf("Test size: ")
	fmt.Scanf("%d", &testSize)
	benchmark(bench.BubbleSort())
	benchmark(bench.InsertionSort())
	benchmark(bench.SelectionSort())
	benchmark(bench.QuickSort())
}

func benchmark(function func(data []int) []int) {

	var elapsed time.Duration
	for i := 1; i <= testSize; i++ {
		array := utils.SortedArray(arraySize, int64(i*testSize*arraySize))
		start := time.Now()
		function(array)
		elapsed += time.Since(start)
	}

	name := getFunctionName(function)
	firstIndex := strings.Index(name, ".")
	lastIndex := strings.Index(name, "Sort")
	name = name[firstIndex+1 : lastIndex]
	fmt.Printf("\n%v sorting...\n", name)

	fmt.Printf("\tSorting time of %v arrays of size %v is %s\n", testSize, arraySize, time.Duration(elapsed))
	fmt.Printf("\tAverage sorting time of %v arrays of size %v is %s\n", testSize, arraySize, time.Duration(elapsed)/(time.Duration(testSize)))

}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
