package main

import (
	"091123/domain"
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
	restaurant := domain.CreateRestaurant("Restaurant", 100, 169)
	fmt.Println("Name: " + restaurant.GetName())
	fmt.Printf("Capacity: %d\n", restaurant.GetCapacity())
	fmt.Printf("Price: %dKč\n", restaurant.GetPrice())

	user := domain.CreateUser("Test", "User", 14, 320)
	fmt.Printf("Balance: %dKč\n", user.GetBalance())
	user.AddMoney(100)
	fmt.Printf("Balance: %dKč\n", user.GetBalance())
}
