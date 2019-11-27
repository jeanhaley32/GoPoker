package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	// Deck  contains a list of cards.
	Deck struct {
		Cards []*Card
	}
	// Card represents an individual playing card.
	Card struct {
		suite     string
		number    int
		character string
		flipped   bool
	}
)

var (
	suites     = [4]string{"diamonds", "spades", "clubs", "hearts"}
	numbers    = [9]int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	characters = [4]string{"jack", "king", "queen", "ace"}
)

// GenDeck method is used to generate a full deck of cards.
func GenDeck() *Deck {
	var d Deck
	for _, n := range suites {
		for _, c := range numbers {
			newCard := Card{suite: n, number: c}
			d.Cards = append(d.Cards, &newCard)
		}
		for _, c := range characters {
			newCard := Card{suite: n, character: c}
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

// Hand returns a hand of cards, with card count  of variable h, removing each card it returns
// from Deck.
func (d *Deck) Hand(h int) ([]*Card, error) {

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
func (d *Deck) ForEachCard(action func(c *Card) error) error {
	for _, card := range d.Cards {
		if err := action(card); err != nil {
			return err
		}
	}
	return nil
}

// ********** CARD METHODS***********
// Read method returns the card type as a string.
func (c *Card) Read() string {
	if c.character != "" {
		return fmt.Sprintf("%v of %v", c.character, c.suite)
	}
	return fmt.Sprintf("%v of %v", c.number, c.suite)
}

// DisplayCards Displays up to ten cards in a row.
func DisplayCards(c []*Card) error {
	if len(c) > 10 {
		return fmt.Errorf("Too many cards to display(%v)", len(c))
	}
	for range c {
		fmt.Printf(" ******** ")
	}
	fmt.Println()

	for _, card := range c {
		fmt.Printf(" * %s    * ", string(card.suite[0]))
	}
	fmt.Println()

	for _, card := range c {
		switch {
		case card.number == 0:
			fmt.Printf(" *  %v   * ", string(card.character[0]))
		case card.number > 9:
			fmt.Printf(" *  %v  * ", card.number)
		default:
			fmt.Printf(" *  %v   * ", card.number)
		}
	}

	fmt.Println()
	for _, card := range c {
		fmt.Printf(" *    %s * ", string(card.suite[0]))
	}
	fmt.Println()
	for range c {
		fmt.Printf(" ******** ")
	}
	fmt.Println()

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
