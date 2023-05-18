package CardGame

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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

func Deal(d Deck, handsize int) (Deck, Deck) {
	return d[:handsize], d[handsize:]
}

func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	n := len(d)
	for i := 0; i < n; i++ {
		j := r.Intn(n)
		d[i], d[j] = d[j], d[i]
	}
}

func (d Deck) ToString() string {
	return strings.Join([]string(d), ",")
}
