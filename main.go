
package main

import (
	"fmt"
	"math/rand"
)
type Card struct {
	Type string
	Suit string
}

type Deck []Card
func New() (deck Deck) {
	types := []string{"Spades","Diamonds","Hearts"}
	suits := []string{"Ace","Two","Three"}

	for i := 0; i < len(types); i++ {
		for n := 0; n < len(suits); n++ {
			card := Card{
				Type: types[i],
				Suit: suits[n],
			}
			deck = append(deck, card)
		}
	}
	return
}

func Shuffle(d Deck) Deck {
	for i := 1; i < len(d); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			d[r], d[i] = d[i], d[r]
		}
	}
	return d
}



func main() {
	deck := New()
	fmt.Printf(" After Append Deck data : - ",deck)
	deck1 :=Shuffle(deck)
   fmt.Printf(" Shuffle Data*********** ",deck1)
}
