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
		Cards  []*Card
		isDeck bool
	}
	// Card represents an individual card
	Card struct {
		Suite     string
		Rank      int
		isFlipped bool
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

// TO-DO; reconsider how this works. it may not be the best method for doing this. Maybe
// utilizing interfaces would be a more robust solution.

// MakeDeck turns a CardCollection into a Deck via switching the isDeck bool
// to true. This allows you to run genDeck.
func (c *CardCollection) MakeDeck() {
	c.isDeck = true
}

// DealCard Methods deals one card form the top of the deck, then deletes that card from the deck.
func (c *CardCollection) DealCard() (Card, error) {
	newCard := new(Card)
	if len(c.Cards) < 1 {
		return *newCard, fmt.Errorf("Cannot return a card from an empty deck")
	}
	newCard = c.Cards[0]
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
		hand.Cards = append(hand.Cards, &newCard)
	}
	return hand, nil
}

// ForEachCard allows an action to be executed on each card in a Deck.
// I have not utilized this method yet, and in the future I may remove it.
func (c *CardCollection) ForEachCard(action func(c *Card) error) error {
	for _, card := range c.Cards {
		if err := action(card); err != nil {
			return err
		}
	}
	return nil
}

// Shuffle func shuffles a deck of cards.
func (c *CardCollection) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for range c.Cards {
		rand.Shuffle(len(c.Cards), func(i, j int) {
			c.Cards[i], c.Cards[j] = c.Cards[j], c.Cards[i]
		})
	}
}

// Read method returns the card type as a string.
func (c *Card) Read() string {
	return fmt.Sprintf("%v of %v", cardIndex[c.Rank], c.Suite)
}

// DisplayCards displays cards in players hand as ascii representations.
func (c *CardCollection) DisplayCards() error {
	fmt.Println("check")
	if len(c.Cards) > 10 {
		return fmt.Errorf("Too many cards to display(%v)", len(c.Cards))
	}
	for i := 0; i < 5; i++ {
		for _, card := range c.Cards {
			num := strconv.FormatInt(int64(card.Rank), 10)
			suit := string(card.Suite[0])
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

// FlipCards flips all cards in a players deck.
// will flip all cards based on the inverse of the
// flipped status of the first card.
func (c *CardCollection) FlipCards() {
	for _, card := range c.Cards {
		card.isFlipped = false
	}
}
