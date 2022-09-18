package main

import "fmt"

func main() {
	i := int(42)
	fmt.Printf("type= %T address=%p value=%v\n", i, &i, i) // type= int address=0xc0000140b0 value=42
}
