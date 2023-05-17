package CardGame

import "fmt"

type Deck []string

func NewDeck() Deck {
	cards := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
		"Nine", "Ten", "Jack", "Queen", "King"}
	suitmarks := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	d := []string{}
	for _, card := range cards {
		for _, suit := range suitmarks {
			d = append(d, card+" of "+suit)
		}
	}
	return d
}

func (d Deck) Show() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
