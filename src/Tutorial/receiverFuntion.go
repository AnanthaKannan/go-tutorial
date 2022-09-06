package main

import (
	"fmt"
	"strings"
)

// to run the file
// go run receiverFunction.go deck.go
func main() {
	// cards := []string{"Ace of diamond", newCard()}
	cards := deck{"Ace of diamond", newCard()}

	cards = append(cards, "Six of spades")
	cards.print()
	fmt.Println("receiverFunction =>>>>>>> ", strings.Join(cards, ", "))
}

func newCard() string {
	return "five of diamonds"
}
