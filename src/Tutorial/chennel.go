package main

import (
	"fmt"
	"time"
)

func portal1(channel1 chan string) {
	time.Sleep(2 * time.Second)
	channel1 <- "Welcome to channel 1"
}

func portal2(channel2 chan string) {
	time.Sleep(2 * time.Second)
	channel2 <- "Welcome to channel 2"
}

func main() {
	R1 := make(chan string)
	R2 := make(chan string)

	// go portal1(R1)
	go portal2(R2)

	select {
	case opt1 := <-R1:
		fmt.Println(opt1)
	case opt2 := <-R2:
		fmt.Println(opt2)
	}
}
