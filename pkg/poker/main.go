package main

import (
	"bufio"
	"deck"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	introMessage = `
	 ****************************
	 *                  ** **   *
	 *                 *  *  *  *
	 *                  *   *   *
	 *                    *     *
	 *                          *
	 *                          *
	 *                          *
	 *                          *
	 *                          *
	 *                          *
	 *                          *
	 *   ** **                  *
	 *  *  *  *                 *
	 *   *   *                  *
	 *     *                    *
	 *                          *
	 ****************************
         Jean's Poker Game
`
)

var (
	handCount = flag.Int("hand_count", 2, "Number of cards in a single hand. ")
)

func main() {

	dealer, _ := initDealer()
	var players []*deck.Player
	players = append(players, dealer.Players...)
	// for _, card := range players[0].Dealer.Deck.Cards {
	// 	fmt.Println(card.Read())
	// }

	for _, player := range players {
		player.Hand.FlipCards(false)
		// player.Hand.DisplayCards()
		for _, card := range player.Hand.Cards {
			fmt.Printf("%v\n", card.Name)
			card.DisplayCard()
		}

		player.FindPairs()
		fmt.Printf("Matches for player %v:\n", player.Name)
		for _, match := range player.HandMatches {
			fmt.Printf("\t%v\n", match.Name)
		}
	}
}

// initDealer initializes the game.
// calls the splash screen.
// prompts for player count.
// initializes a dealer with players, and deck.
func initDealer() (deck.Dealer, error) {
	splashScreen()
	var dealer deck.Dealer
	dealer.InitDeck()
	pCount := 0
	for {
		splashScreen()
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter up to 5 players.")
		sBucket, err := reader.ReadString('\n')
		if err != nil {
			return dealer, fmt.Errorf("Failed to read from stdin:%v", err)
		}
		pCount, err = strconv.Atoi(sBucket[:len(sBucket)-1])
		if err != nil {
			fmt.Println("Please provide a number.")
			pause(2)
			continue
		}
		if pCount <= 10 {
			break
		} else {
			fmt.Println("Max number of players is nine.")
			pause(1)
			continue
		}

	}
	for i := 0; i < pCount; i++ {
		splashScreen()
		fmt.Printf("%v ", i+1)
		_ = dealer.GenPlayer()
	}
	clear()
	return dealer, nil
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func splashScreen() {
	clear()
	fmt.Println(introMessage)
}

func pause(i time.Duration) {
	time.Sleep(i * time.Second)
}
