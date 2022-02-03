package main

import "fmt"

func main() {
	msg := "Hello world"
	writingMessage(msg)
	count := sum("", 1, 2, 3, 4, 5)
	fmt.Println("count-1 ", count)
	count = sum("", 10, 20, 11, 24, 11, 15)
	fmt.Println("count-2 ", count)

	value, err := div(1, 20)
	if err == nil {
		fmt.Println("value-div", value)
	}
}

func writingMessage(msg string) {
	fmt.Println(msg)
}

// This function called as variadic functions
// values ...int this is variadic
func sum(msg string, values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	// fmt.Println("sumTotal", total)
	return total
}

func div(a float64, b float64) (float64, error) {

	if b == 0.0 || b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	result := a / b
	return result, nil
}
