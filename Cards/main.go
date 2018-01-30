package main

import (
	"fmt"
)

func main() {

	//var card = "Ace of spades"
	//card := "Ace of Spades"
	//card = "Five of Diamond" //':' not required after first initialization.
	//card := newCard()

	//Creating slices in Go (A type of array that can grow and shrink)
	cardArray := []string{"Ace of Diamonds", newCard(), newCard()} //A slice of type string

	cardArray = append(cardArray, "Six of Spades") // how to append a "slice"or array

	fmt.Println(cardArray)

	//Iterating over a slice
	for i, card := range cardArray {
		fmt.Println(card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
