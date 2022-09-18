package main

import "fmt"

func main() {
	result := returnInterface(1)
	fmt.Println(result)
}

// if the interface means it return anything, it can be number, string or anything
func returnInterface(number int) interface{} {
	if number == 1 {
		return true
	}
	if number == 2 {
		return "some string"
	}
	return nil
}
