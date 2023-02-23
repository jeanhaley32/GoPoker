package deck

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type (
	// CardCollection represents any list of cards.
	// I.E. deck, hand etc.
	CardCollection struct {
		Cards []Card
	}
	// Card represents an individual card
	Card struct {
		Name      string
		Suite     int
		Rank      int
		isFlipped bool
	}
)

var (
	suites    = [4]int{0, 1, 2, 3}
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
	suitIndex = map[int]string{
		0: "clubs",
		1: "diamonds",
		2: "hearts",
		3: "spades",
	}
)

// GenDeck returns an un-shuffled deck of cards.
func (c *CardCollection) GenDeck(shuffle bool) CardCollection {
	var deck CardCollection
	for _, s := range suites {
		for _, r := range ranks {
			newCard := Card{
				Name:  fmt.Sprintf("%v of %v", cardIndex[r], suitIndex[s]),
				Suite: s,
				Rank:  r,
			}
			newCard.isFlipped = true
			deck.Cards = append(deck.Cards, newCard)
		}
		if shuffle == true {
			deck.Shuffle()
			deck.Shuffle()
		}
	}
	return deck
}

// DealCard Methods deals one card form the top of the deck, then deletes that card from the deck.
func (c *CardCollection) DealCard() (Card, error) {
	newCard := new(Card)
	if len(c.Cards) < 1 {
		return *newCard, fmt.Errorf("Cannot return a card from an empty deck")
	}
	newCard = &c.Cards[0]
	c.removeCard(0)
	return *newCard, nil
}

func (c *CardCollection) removeCard(i int) error {
	cardCount := len(c.Cards)
	if cardCount > 0 {
		c.Cards[i] = c.Cards[cardCount-1]
		c.Cards = c.Cards[:cardCount-1]
		return nil
	}
	return fmt.Errorf("cannot remove card(card count %v", cardCount)
}

// Count returns the amount of cards in a deck of cards.
func (c *CardCollection) Count() int {
	return len(c.Cards)
}

// DealCards returns a hand of cards, with card count  of variable h, removing each card it returns
// from Deck.
func (c *CardCollection) DealCards(h int) (CardCollection, error) {
	var hand CardCollection
	for i := 0; i <= h-1; i++ {
		newCard, err := c.DealCard()
		if err != nil {
			return hand, fmt.Errorf("Can't provide a hand of cards %v", err)
		}
		hand.Cards = append(hand.Cards, newCard)
	}
	return hand, nil
}

// ForEachCard allows an action to be executed on each card in a Deck.
// I have not utilized this method yet, and in the future I may remove it.
func (c *CardCollection) ForEachCard(action func(c *Card) error) error {
	for _, card := range c.Cards {
		if err := action(&card); err != nil {
			return err
		}
	}
	return nil
}

// Shuffle randomizes the order of cards within a CardCollection.
func (c *CardCollection) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for range c.Cards {
		rand.Shuffle(len(c.Cards), func(i, j int) {
			c.Cards[i], c.Cards[j] = c.Cards[j], c.Cards[i]
		})
	}
}

// // Read method returns the card type as a string.
// // rethink the need for this, now that I have each card contain it's own canonical name.
// func (c *Card) Read() string {
// 	return fmt.Sprintf("%v of %v", cardIndex[c.Rank], c.Suite)
// }

// DisplayCards displays each card in players hand as ascii art representations, in a row.
func (c *CardCollection) DisplayCards() error {
	if len(c.Cards) > 10 {
		return fmt.Errorf("Too many cards to display(%v)", len(c.Cards))
	}
	for i := 0; i < 5; i++ {
		for _, card := range c.Cards {
			num := strconv.FormatInt(int64(card.Rank), 10)
			suit := string(suitIndex[card.Suite][0])
			rank := string(num) + " "

			switch {
			case card.isFlipped == true:
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

// DisplayCard displays an ascii Representation of one card.
func (c *Card) DisplayCard() error {
	for i := 0; i < 5; i++ {

		num := strconv.FormatInt(int64(c.Rank), 10)
		suit := string(suitIndex[c.Suite][0])
		rank := string(num) + " "

		switch {
		case c.isFlipped == true:
			rank = "  "
			suit = "?"
		case c.Rank >= 11:
			rank = string(cardIndex[c.Rank][0]) + " "
		case c.Rank == 10:
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
		fmt.Println()
	}
	return nil
}

// FlipCards sets each card in a card collection to
// isFlipped true/false based on input.
//
// true will flip cards, so when the ascii value is displayed, it will
// show as a card with question marks.
func (c *CardCollection) FlipCards(d bool) {
	for i := range c.Cards {
		if d == false {
			c.Cards[i].isFlipped = false
		} else {
			c.Cards[i].isFlipped = true
		}
	}
}
