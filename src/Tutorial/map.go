package main

import "fmt"

func main() {
	// key will be string and value will be int
	shoppingCard := map[string]int{
		"keyboard": 100,
		"mouse":    50,
		"cpu":      201,
	}
	shoppingCard["monitor"] = 121 // add the data

	delete(shoppingCard, "cpu") // delete the data

	fmt.Println("shoppingCard", shoppingCard, "shoppingCard_len", len(shoppingCard))
	fmt.Println("keyboardAmount", shoppingCard["keyboard"])

	// if you want to check the value is available or no
	_, ok := shoppingCard["mobile"]
	fmt.Println("isAvailable", ok)
}
