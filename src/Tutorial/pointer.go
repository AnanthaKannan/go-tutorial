package main

import "fmt"

// pointer used to pointing the memory address
func main() {
	var a int = 20
	var b *int = &a
	// in this case it is not copy the value, it is just copy the memory address from a to b
	fmt.Println(a, b, *b)
}
