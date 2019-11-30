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

// DisplayCards displays cards in players hand as ascii representations.
func (p *Player) DisplayCards() error {
	if len(p.Hand) > 10 {
		return fmt.Errorf("Too many cards to display(%v)", len(p.Hand))
	}
	for i := 0; i < 5; i++ {
		for _, card := range p.Hand {
			rank := string(cardIndex[card.rank][0])
			suit := string(card.suite[0])
			switch {
			case i == 0:
				fmt.Printf(" ******** ")
			case i == 1:
				fmt.Printf(" * %v    * ", rank)
			case i == 2:
				fmt.Printf(" *   %v  * ", suit)
			case i == 3:
				fmt.Printf(" *    %v * ", rank)
			case i == 4:
				fmt.Printf(" ******** ")
			}
		}
		fmt.Println()
	}
	return nil
}
