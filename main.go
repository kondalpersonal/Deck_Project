package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamons", newCard()}
	cards = append(cards, "Three of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Two of Loves"
}
