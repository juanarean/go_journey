package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // como no necesito el index lo llamo _ y go entiende que no necesito esa variable que devuelve range
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] 
}