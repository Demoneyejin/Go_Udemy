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
	/*cardArray := []string{"Ace of Diamonds", newCard(), newCard()} //A slice of type string

	cardArray = append(cardArray, "Six of Spades") // how to append a "slice"or array

	fmt.Println(cardArray)

	//Iterating over a slice
	for i, card := range cardArray {
		fmt.Println(card)
	}
	*/
	//All of the commented code above was done via previous udemy learning. Now we will be building
	//The Card Example.

	cards := deck{"Ace of Diamonds", newCard()}

	fmt.Println(cards) //used to avoid error lines.

}

func newCard() string {
	return "Five of Diamonds"
}
