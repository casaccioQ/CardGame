package CardGame

import "testing"

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	n := len(d)
	if n != 52 {
		t.Errorf("Expected deck length of 52, but got %v.\n", n)
	}

	first := d[0]
	if first != "Ace of Spades" {
		t.Errorf("The first card in the deck should be the Ace of Spades, but it is %v instead.\n", first)
	}

	last := d[n-1]
	if last != "King of Diamonds" {
		t.Errorf("The last card in the deck should be the King of Diamonds, but it is %v instead.\n", last)
	}
}
