package main

import (
	"ProjektBezeJmena/domain"
	"fmt"
	// "strconv"
)

// func main() {
// 	// y := 2
// 	// p := &y
// 	// fmt.Println(p)
// 	// fmt.Println(*p)

// 	number := "123456789"
// 	parsedNumber, err := ParseString(number)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}
// 	fmt.Println("Number: ", parsedNumber)
// }

// func ParseString(s string) (*int64, error) {
// 	parsed, err := strconv.ParseInt(s, 10, 64)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &parsed, nil
// }


func main() {
	client := domain.CreateUser("Richard", "Kazbundič", 93, 100)
	fmt.Printf("Client: %+v\n", client)
}