package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	var i int = 101 // you can mention type as well
	i = 102
	var name = "Kannan"
	fmt.Println(name)
	fmt.Printf("%v, %T\n", i, i)

	// another way to declare the variable
	age := 22 // it will declare the variable and assign the value
	print("age ", age)
}
