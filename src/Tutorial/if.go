package main

import "fmt"

func main() {
	if true {
		fmt.Println("it is inside the if condition")
	}

	shoppingCard := map[string]int{
		"keyboard": 100,
		"mouse":    50,
		"cpu":      202,
	}

	cpu := shoppingCard["cpu"]
	if cpu == 201 {
		fmt.Println("It is valid")
	} else if cpu == 202 {
		fmt.Println("It is inside the else if")
	} else {
		fmt.Println("it is inside the else")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i ", i)
	}

	arr := []int{10, 11, 12, 13, 14, 15, 16}
	for k, v := range arr {
		fmt.Println(k, v)
	}
}
