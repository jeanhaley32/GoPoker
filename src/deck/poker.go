package deck

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// Hand represents an individual hand, and its value.
	Hand struct {
		name  string
		value int
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

// GetHands take in a list of cards, and figureds out what hands exist within that list, 
// return as many matches as it finds as a set fo []hands, each hand contains the name 
// of the matching hand and that hands value.
func GetHands(c []*Card) ([]Hand, error) {

	// sort references
	//sort.Slice(c, func(i, j int) bool { return c[i].Suite < c[j].Suite })
	//sort.Slice(c, func(i, j int) bool { return c[i].Rank < c[j].Rank })

	// straight flush
	// four of a kind
	// full house
	// flush
	// straight
	// three of a kind
	// two pair
	// one pair
	// high card
	var hands []Hand
	suitSorted := map[string][]int{"clubs": {}, "diamonds": {}, "hearts": {}, "spades": {}}

	royalFlush := []int{14, 13, 12, 11, 10}

	for _, card := range c {
		switch {
		case card.Suite == "clubs":
			suitSorted[card.Suite] = append(suitSorted[card.Suite], card.Rank)
		case card.Suite == "diamonds":
			suitSorted[card.Suite] = append(suitSorted[card.Suite], card.Rank)
		case card.Suite == "hearts":
			suitSorted[card.Suite] = append(suitSorted[card.Suite], card.Rank)
		case card.Suite == "spades":
			suitSorted[card.Suite] = append(suitSorted[card.Suite], card.Rank)
		}
	}

	for key, element := range suitSorted {
		sort.Slice(element, func(i, j int) bool { return suitSorted[key][i] < suitSorted[key][j] })

	}

	for key, element := range suitSorted {
		for _, r := range element {
			fmt.Printf("%v of %v\n", cardIndex[r], key)
		}
	}

	return nil, nil
}
