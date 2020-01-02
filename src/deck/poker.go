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
		Name      string
		Value     int
		Suit      string
		HighCard  int
		PairCount int
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

// InitDeck Initializes the dealers deck.
func (d *Dealer) InitDeck() {
	d.Deck = d.Deck.GenDeck(true)
}

// GenPlayer generates a new player object tied to a dealer.
func (d *Dealer) GenPlayer() *Player {
	var player Player
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Players Name?\n")
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to obtain name: %v\nplease try again.", err)
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

// FindPairs finds two of a kind, three of a kind, four of a kind, and Full House.
func (p *Player) FindPairs() {
	pHand := p.Hand
	dHand := p.Dealer.Table
	testHand := append(pHand.Cards, dHand.Cards...)
	rankMap := make(map[int]int)
	var matches []HandMatch
	for _, card := range testHand {
		rankMap[card.Rank]++
	}
	for rank, count := range rankMap {
		switch {
		case count == 2:
			handName := fmt.Sprintf("%v(%v)", twoPair, cardIndex[rank])
			handValue := handValueIndex[twoPair] + rank
			match := HandMatch{
				Name:      handName,
				Value:     handValue,
				HighCard:  rank,
				PairCount: 2,
			}
			matches = append(matches, match)
		case count == 3:
			handName := fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[rank])
			handValue := handValueIndex[threeOfaKind] + rank
			match := HandMatch{
				Name:      handName,
				Value:     handValue,
				HighCard:  rank,
				PairCount: 3,
			}
			matches = append(matches, match)
		case count == 4:
			handName := fmt.Sprintf("%v(%v)", fourOfaKind, cardIndex[rank])
			handValue := handValueIndex[fourOfaKind] + rank
			match := HandMatch{
				Name:      handName,
				Value:     handValue,
				HighCard:  rank,
				PairCount: 4,
			}
			matches = append(matches, match)
		}
	}
	if (len(matches) >= 2) && (matches[0].PairCount+matches[1].PairCount >= 5) {
		fmt.Println("potential FullHouse Match")
		switch {
		case len(matches) == 2 && matches[0].PairCount == 2 || matches[1].PairCount == 2:
			highestValue := 0
			fmt.Println("Got em")
			for _, match := range matches {
				if match.PairCount == 3 || match.PairCount == 4 {
					highestValue = match.PairCount
				}
			}
			match := HandMatch{
				Name:     fmt.Sprintf("%v(%v)", fullHouse, highestValue),
				HighCard: highestValue,
				Value:    handValueIndex[fullHouse] + highestValue,
			}
			matches = append(matches, match)
		default:
			fmt.Println("not a FullHouse")
		}
	}
	p.HandMatches = append(p.HandMatches, matches...)
}

// suitSorted returns a map of suit sorted cards from the players hand.
func (p *Player) suitSorted() map[int][]int {

	suitSorted := map[int][]int{0: {}, 1: {}, 2: {}, 3: {}}

	for _, card := range p.Hand.Cards {
		switch {
		case card.Suite == 0:
			suitSorted[card.Suite] = append(suitSorted[card.Suite], card.Rank)
		case card.Suite == 1:
			suitSorted[card.Suite] = append(suitSorted[card.Suite], card.Rank)
		case card.Suite == 2:
			suitSorted[card.Suite] = append(suitSorted[card.Suite], card.Rank)
		case card.Suite == 3:
			suitSorted[card.Suite] = append(suitSorted[card.Suite], card.Rank)
		}
	}

	for key, element := range suitSorted {
		sort.Slice(element, func(i, j int) bool { return suitSorted[key][i] < suitSorted[key][j] })
	}

	return suitSorted
}
