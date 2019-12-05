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
		Name   string
		Hand   *CardCollection
		Dealer *Dealer
	}

	// Dealer represents the dealer, with a deck of cards, and the table.
	Dealer struct {
		Deck    *CardCollection
		Players []*Player
		Table   []Card
	}

	// Hand represents an individual hand, and its value.
	Hand struct {
		name     string
		value    int
		suit     string
		highCard int
	}
)

// GenDealer generates a new Dealer with a fresh deck of cards.
func GenDealer() Dealer {
	var newDealer *Dealer
	newDealer.Deck = GenDeck()
	newDealer.Deck.Shuffle()
	return *newDealer
}

// GenPlayer generates a new player object tied to a dealer.
func (d *Dealer) GenPlayer() {
	var player *Player
	player.Dealer = d
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Players Name?\n")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to obtain name: %w\nplease try again.", err)
			continue
		} else {
			player.Name = name[:len(name)-2]
			// TO-DO strip spaces from the end of this^^
			break
		}
	}
	d.Players = append(d.Players, player)

}

// FlipCards flips all cards in a players deck.
// will flip all cards based on the inverse of the
// flipped status of the first card.
func (p *Player) FlipCards() bool {
	var flip bool
	if p.Hand.Cards[0].Flipped == false {
		flip = true
	}
	for _, card := range p.Hand.Cards {
		card.Flipped = flip
	}
	return flip

}

// GetHands take in a list of cards, and figureds out what hands exist within that list,
// return as many matches as it finds as a set fo []hands, each hand contains the name
// of the matching hand and that hands value.
func GetHands(c []*Card) ([]Hand, error) {

	// sort references
	//sort.Slice(c, func(i, j int) bool { return c[i].Suite < c[j].Suite })
	//sort.Slice(c, func(i, j int) bool { return c[i].Rank < c[j].Rank })

	handValue := map[string]int{
		"royal flush":     10,
		"straight flush":  9,
		"four of a kind":  8,
		"full house":      7,
		"flush":           6,
		"straight":        5,
		"three of a kind": 4,
		"two pair":        3,
		"one pair":        2,
		"high card":       1,
	}

	var hands []Hand
	suitSorted := map[string][]int{"clubs": {}, "diamonds": {}, "hearts": {}, "spades": {}}

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

	for suit, ranks := range suitSorted {

		testCase := ranks
		if len(ranks) > 10 {
			testCase = testCase[0:5]
		}

		for i := len(testCase); i > 0; i-- {

		}

		switch {
		case royalFlush(ranks):
			hands = append(hands, Hand{
				name:     "royal flush",
				value:    handValue["royal flush"],
				suit:     suit,
				highCard: ranks[len(ranks)-1],
			})
		}

		for key, element := range suitSorted {
			for _, r := range element {
				fmt.Printf("%v of %v\n", cardIndex[r], key)
			}
		}
	}
	return nil, nil
}

func royalFlush(c []int) bool {
	royalFlush := []int{10, 11, 12, 13, 14}
	length := len(c)
	if len(c) > 5 {
		c = c[length-5:]
	}
	if Equal(c, royalFlush) {
		return true
	}
	return false
}

// func flush(c []int) bool {
// 	// function to detect a flush.
// 	length := len(c)
// 	switch {
// 	case length > 5:
// 		front := c[0:5]
// 		tail := c[length-5:]
// 	}
// 	return false
// }

// Equal tells whether presorted lists a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
