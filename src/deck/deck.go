package deck

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type (
	// Deck  contains a list of cards.
	Deck struct {
		Cards []*Card
	}
	// Card represents an individual playing card.
	Card struct {
		Suite   string
		Rank    int
		Flipped bool
	}
)

var (
	suites    = [4]string{"diamonds", "spades", "clubs", "hearts"}
	ranks     = [13]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	cardIndex = map[int]string{
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "jack",
		12: "queen",
		13: "king",
		14: "ace",
	}
)

// GenDeck method is used to generate a full deck of cards.
func GenDeck() *Deck {
	var d Deck
	for _, n := range suites {
		for _, c := range ranks {
			newCard := Card{Suite: n, Rank: c}
			newCard.Flipped = true
			d.Cards = append(d.Cards, &newCard)
		}
	}
	return &d
}

// Card Method pulls a card from a deck of cards.
// it will also remove this card from the deck.
func (d *Deck) Card() (Card, error) {
	newCard := new(Card)
	if len(d.Cards) < 1 {
		return *newCard, fmt.Errorf("Cannot return a card from an empty deck")
	}
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Intn(len(d.Cards))
	index -= index

	*newCard = *d.Cards[index]
	d.removeCard(index)
	return *newCard, nil
}

func (d *Deck) removeCard(i int) error {
	cardCount := len(d.Cards)
	if cardCount > 0 {
		d.Cards[i] = d.Cards[cardCount-1]
		d.Cards = d.Cards[:cardCount-1]
		return nil
	}
	return fmt.Errorf("cannot remove card(card count %v", cardCount)
}

// Count returns the amount of cards in a deck of cards.
func (d *Deck) Count() int {
	return len(d.Cards)
}

// GetHand returns a hand of cards, with card count  of variable h, removing each card it returns
// from Deck.
func (d *Deck) GetHand(h int) ([]*Card, error) {
	var hand []*Card
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i <= h-1; i++ {
		newCard, err := d.Card()
		if err != nil {
			return hand, fmt.Errorf("Can't provide a hand of cards %v", err)
		}
		hand = append(hand, &newCard)
	}
	return hand, nil
}

// ForEachCard allows an action to be executed on each card in a Deck.
// I have not utilized this method yet, and in the future I may remove it. 
func (d *Deck) ForEachCard(action func(c *Card) error) error {
	for _, card := range d.Cards {
		if err := action(card); err != nil {
			return err
		}
	}
	return nil
}

// Shuffle func shuffles a deck of cards.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for range d.Cards {
		rand.Shuffle(len(d.Cards), func(i, j int) {
			d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
		})
	}
}

// Read method returns the card type as a string.
func (c *Card) Read() string {
	return fmt.Sprintf("%v of %v", cardIndex[c.Rank], c.Suite)
}

// DisplayCards displays cards in players hand as ascii representations.
func DisplayCards(c []*Card) error {
	if len(c) > 10 {
		return fmt.Errorf("Too many cards to display(%v)", len(c))
	}
	for i := 0; i < 5; i++ {
		for _, card := range c {
			num := strconv.FormatInt(int64(card.Rank), 10)
			suit := string(card.Suite[0])
			rank := string(num) + " "

			switch {
			case card.Flipped == true:
				rank = "  "
				suit = "?"
			case card.Rank >= 11:
				rank = string(cardIndex[card.Rank][0]) + " "
			case card.Rank == 10:
				rank = "10"
			}

			switch {
			case i == 0:
				fmt.Printf(" ******** ")
			case i == 1:
				fmt.Printf(" * %v   * ", rank)
			case i == 2:
				fmt.Printf(" *   %v  * ", suit)
			case i == 3:
				fmt.Printf(" *    %v* ", rank)
			case i == 4:
				fmt.Printf(" ******** ")
			}
		}
		fmt.Println()
	}
	return nil
}
