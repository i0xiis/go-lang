package main

import (
	"fmt"
)

//* import "fmt"

func main() {
	list := []int{5,8,3,2,9,7,4,1,6,10}
	for i := 0; i < len(list); i++ {
		bigIdx := -1
		for j := 0; j < len(list) - i; j++ {
			if bigIdx == -1 || list[j] > list[bigIdx] {
				bigIdx = j
			}
		}
		Swap(&list[bigIdx], &list[len(list) - i - 1])
	}
	// a := 10
	// b := -10
	// Swap(&a, &b)
	// fmt.Println(a, b)

	fmt.Println(list)
}

func Swap(x, y *int) {
	c := *y
	*y = *x
	*x = c
}