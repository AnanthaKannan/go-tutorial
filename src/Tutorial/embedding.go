package main

import "fmt"

// processor struct
type Processor struct {
	processorName string
	cores         int
}

// memory struct
type Memory struct {
	memoryCapacity int
	MemoryType     string
}

// Computer struct
type Computer struct {
	Processor
	Memory
	price int
}

func main() {
	// one way of value assigning
	computerA := Computer{}
	computerA.price = 54000
	computerA.processorName = "INTEL"

	// another way of value assigning
	computerB := Computer{
		Processor: Processor{
			processorName: "RAZAON",
			cores:         4,
		},
		Memory: Memory{
			memoryCapacity: 5,
		},
	}
	fmt.Println(computerB)
}
