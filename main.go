package main

import "fmt"

func main() {
	fmt.Println("Welcome to Blackjack. Basic rules follow. ")
	cards := createDeck()
	cards = cards.shuffle()
	cards.deal()
	cards.gameStart()
}
