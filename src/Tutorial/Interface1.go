package main

import "fmt"

func main() {
	// fmt.Println("Hello World 1")
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World 2"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Writer(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
