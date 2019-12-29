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
		Name        string
		Hand        CardCollection
		HandMatches []HandMatch
		Dealer      *Dealer
	}

	// Dealer represents the dealer, with a deck of cards, and the table.
	Dealer struct {
		Deck    CardCollection
		Players []*Player
		Table   CardCollection
	}

	// HandMatch represents an individual hand, and its value.
	HandMatch struct {
		name     string
		value    int
		suit     string
		highCard int
	}
)

var (
	handValueIndex = make(map[string]int)

	// Creating referenceable strings, to make naming universal.
	royalFlush    = "royal flush"
	straightFlush = "straight flush"
	fourOfaKind   = "four of a kind"
	fullHouse     = "full house"
	flush         = "flush"
	straight      = "straight"
	threeOfaKind  = "three of a kind"
	twoPair       = "two pair"
	onePair       = "one pair"
	highCard      = "high card"
)

func init() {
	handValueIndex = map[string]int{
		royalFlush:    10,
		straightFlush: 9,
		fourOfaKind:   8,
		fullHouse:     7,
		flush:         6,
		straight:      5,
		threeOfaKind:  4,
		twoPair:       3,
		onePair:       2,
		highCard:      1,
	}
}

// GenDeck populates the dealers deck of cards.
// TO-DO: split this into two functions, the deck should
// generate from the deck library, and then this function
// should only deal with appending that deck to the dealer, and maybe shuffling it.
// consider adding shuffling as an optional variable passed to the genDeck function, once
// it's located within deck.go. Maybe rename this one to genDealerDeck
func (d *Dealer) GenDeck() {
	for _, s := range suites {
		for _, r := range ranks {
			var newCard Card
			newCard.Suite = s
			newCard.Rank = r
			newCard.isFlipped = true
			d.Deck.Cards = append(d.Deck.Cards, &newCard)
		}
		d.Deck.Shuffle()
	}
}

// GenPlayer generates a new player object tied to a dealer.
func (d *Dealer) GenPlayer() *Player {
	var player Player
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Players Name?\n")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to obtain name: %w\nplease try again.", err)
			continue
		} else {
			player.Name = name[:len(name)-1]
			// TO-DO strip spaces from the end of this^^
			break
		}
	}
	player.Dealer = d
	player.Hand, _ = d.Deck.DealCards(5)

	d.Players = append(d.Players, &player)
	return &player
}

// func GetHandMatches(p *Player) ([]HandMatch, error) {

// 	// sort references
// 	//sort.Slice(c, func(i, j int) bool { return c[i].Suite < c[j].Suite })
// 	//sort.Slice(c, func(i, j int) bool { return c[i].Rank < c[j].Rank })

// // func (p *Player)royalFlush()  {
// // 	royalFlush := []int{10, 11, 12, 13, 14}

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

// FindPairs detects duplicate cards within a players hand.
func (p *Player) FindPairs() {
	pHand := p.Hand
	dHand := p.Dealer.Table
	pHandMatches := p.HandMatches
	testHand := append(pHand.Cards, dHand.Cards...)
	rankMap := make(map[int]int)
	var matches []HandMatch
	for _, card := range testHand {
		rankMap[card.Rank]++
	}
	for key, value := range rankMap {
		switch {
		case value == 2:
			handName := fmt.Sprintf("%v(%v)", twoPair, cardIndex[key])
			handValue := handValueIndex[twoPair] + key
			match := HandMatch{
				name:     handName,
				value:    handValue,
				highCard: key,
			}
			matches = append(matches, match)
		case value == 3:
			handName := fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[key])
			handValue := handValueIndex[twoPair] + key
			match := HandMatch{
				name:     handName,
				value:    handValue,
				highCard: key,
			}
			matches = append(matches, match)
		case value == 4:
			handName := fmt.Sprintf("%v(%v)", fourOfaKind, cardIndex[key])
			handValue := handValueIndex[twoPair] + key
			match := HandMatch{
				name:     handName,
				value:    handValue,
				highCard: key,
			}
			matches = append(matches, match)
		}
	}
	pHandMatches = append(pHandMatches, matches...)
}

// suitSorted returns a map of suit sorted cards from the players hand.
func (p *Player) suitSorted() map[string][]int {

	suitSorted := map[string][]int{"clubs": {}, "diamonds": {}, "hearts": {}, "spades": {}}

	for _, card := range p.Hand.Cards {
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

	return suitSorted
}
