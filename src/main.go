// learn to do tests in golang
package main

import (
	"fmt"

	"github.com/saul178/algorithms/src/algorithms"
)

func main() {
	testArray := [5]int{3, 4, 1, 24, 6}
	testLinearSearch := algorithms.BubbleSort(testArray[:])
	fmt.Print(testLinearSearch)
}
