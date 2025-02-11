package main

import (
	"fmt"
	"DeckOfCards/deck"
)

func main() {
	deck := deck.NewDeck()
	deck.Shuffle()
	hand, err := deck.Deal(5)
	if err != nil{
		fmt.Println(err)
	}
	hand.Print()
	fmt.Println("************************")
	deck.Print()
	deck.SaveDeckToFile("my_deck.txt")
}
