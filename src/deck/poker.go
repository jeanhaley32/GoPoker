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
		Hand []*Card
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
	newDealer.Deck.Shuffle()
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

func (p *Player) removeCard(i int) error {
	cardCount := len(p.Hand)
	if cardCount > 0 {
		p.Hand[i] = p.Hand[cardCount-1]
		p.Hand = p.Hand[:cardCount-1]
		return nil
	}
	return fmt.Errorf("cannot remove card(card count %v", cardCount)
}

// FlipCards flips all cards in a players deck.
// will flip all cards based on the inverse of the
// flipped status of the first card.
func (p *Player) FlipCards() bool {
	var flip bool
	if p.Hand[0].Flipped == false {
		flip = true
	}
	for _, card := range p.Hand {
		card.Flipped = flip
	}
	if flip == false {
		return true
	}
	return false

}

// GetHands matches, takes in players hand, and dealers hand, then derives
// a list of valid poker hands from the combination of both.
func GetHands(p, d []*Card) ([]string, error) {

	// royal flush
	cards := append(p, d...)
	for _, card := range cards {
		rank := card.Rank
		suite := card.Suite

	}
	// straight flush
	// four of a kind
	// full house
	// flush
	// straight
	// three of a kind
	// two pair
	// one pair
	// high card
	return nil, nil
}
