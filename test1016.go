package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := make([]int, len(a))
	copy(b, a)
	fmt.Println(b)
	fmt.Println(a)
}
