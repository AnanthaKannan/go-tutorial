package main

import "fmt"

func main() {
	var myInt Incrementor = &IntCounter{no: 10}
	fmt.Println(myInt.Increment())
	fmt.Println(myInt.Subtract())
	// fmt.Println(myInt.Add())
	fmt.Println()
}

type Incrementor interface {
	Increment() int
	Subtract() int
	Add() int
}

type TestInterface interface {
	TestA() string
	TestB() int
}

func TestA() int {
	return 100
}

type IntCounter struct {
	no int
}

func (ic *IntCounter) Increment() int {
	ic.no++
	return ic.no
}

func (ic *IntCounter) Subtract() int {
	fmt.Println("Subtract", ic.no)
	ic.no--
	return ic.no
}

func (ic *IntCounter) Add() int {
	return 10
}
