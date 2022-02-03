package main

import "fmt"

// way to declare multiple constant
const (
	User    string = "Admin"
	Product string = "myProduct"
)

// iota is used to increase the assigned value. example x will be 0 and y will be 1
const (
	x = iota
	y = iota
)

func main() {
	// constant can declare and with out use but variable should use if you are declare
	const i int = 12
	const j float32 = 12.2
	const k string = "this is constant"
	const l bool = true

	fmt.Println(k)
}
