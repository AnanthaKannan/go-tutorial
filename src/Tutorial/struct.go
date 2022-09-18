package main

import "fmt"

// struck is the collection of data type
// we can say another way this as a object

// student struct
type Student struct {
	name     string
	rollNo   int
	subjects []string
}

func main() {
	student1 := Student{
		name:   "Kannan",
		rollNo: 81,
		subjects: []string{
			"maths", "English", "Tamil",
		},
	}

	student1.rollNo = 82
	fmt.Println(student1, student1.name)
}

func anotherFn() {
	println("eeeeeeeeeeeeeeeee")
}
