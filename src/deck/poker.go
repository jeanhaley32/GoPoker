package deck

import (
	"bufio"
	"fmt"
	"os"
)

type (
	// Player represents an individual player with name and a hand of cards.
	Player struct {
		Name string
		Hand []Card
	}

	// Dealer represents the dealer, with a deck of cards, and the table.
	Dealer struct {
		Deck  *Deck
		Table []Card
	}
)

// GenDealer generates a new Dealer with a fresh deck of cards.
func GenDealer() Dealer {
	var newDealer = new(Dealer)
	newDealer.Deck = GenDeck()
	return *newDealer
}

// GenPlayer generates a new player, asking for their name.
func GenPlayer() Player {
	var newPlayer = new(Player)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Players Name?\n")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to obtain name: %w\nplease try again.", err)
			continue
		} else {
			newPlayer.Name = name[:len(name)-1]
			break
		}
	}
	return *newPlayer

}
