package main

import "fmt"

// defer act just like a non blocking, if you are using defer keyword it will run as a last

// if you are using defer for all the lines then it will execute form the last
// act like FILO

func main() {
	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}
