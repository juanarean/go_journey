package main

// import "fmt"

func main() {
	//card := "Ace of Spades" // var card string = "Ace of Spades"
	cards := newDeck() // deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// fmt.Println(cards)

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }