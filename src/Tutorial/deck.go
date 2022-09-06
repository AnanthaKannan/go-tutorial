package main

import (
	"fmt"
	"strings"
)

// Create new type of Deck
// which is type of string
// we can access this deck globally in the main package
type deck []string

/*
we have used receiver function here,
It is means any variable of type "deck" now gets access to the "print" method
*/
func (d deck) print() {
	fmt.Println(d)
	fmt.Println("deckFn==========> ", strings.Join(d, ", "))
}
