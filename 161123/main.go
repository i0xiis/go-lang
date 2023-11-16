package main

import (
	"fmt"
)

//* import "fmt"

func main() {
	list := []int{5,8,3,2,9,7,4,1,6,10}
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d\n", list[i])
	}
}