// learn to do tests in golang
package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a[:len(a)-1])
}
