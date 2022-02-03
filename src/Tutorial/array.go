package main

import "fmt"

func main() {

	// one way to declare the array
	var amounts [3]int = [3]int{10, 20, 30}
	// second way to declare the array
	amt := [3]int{40, 50, 10}
	// if you don't know the array length then you can follow
	am := [...]int{10, 20, 50, 30, 20, 11, 43}

	amounts[1] = 1000

	fmt.Println("amounts: ", amounts, "amt-length", len(amt), am)
	fmt.Println("particular index", amounts[1])

	// get the data
	fmt.Println("print all the data", am[:])
	fmt.Println("print the data from 2nd element", am[2:])
	fmt.Println("print the data up to 5th element", am[:5])
	fmt.Println("print from second element to 3rd element", am[2:4])

	// basically this part is slices. you can create array like this but if you are try to copy array form one to another
	// then its copy the memory address not a value
	var a []int = []int{1, 2, 3}
	var b []int = a
	fmt.Println("value of b", b)

	a = append(a, 10)
	a = append(a, a...) // second argument should be number
	fmt.Println("print the a value", a)
}
