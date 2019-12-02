package main

import (
	"bufio"
	"deck"
	"flag"
	"fmt"
	"log"
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
	handCount = flag.Int("hand_count", 5, "Number of cards in a single hand. ")
)

func main() {
	var err error
	dealer := deck.GenDealer()
	player := deck.GenPlayer()
	player.Hand, err = dealer.Deck.GetHand(5)
	if err != nil {
		fmt.Println(err)
	}
	deck.DisplayCards(player.Hand)

	answer := new(string)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("would you like to reveal your cards?")
		*answer, err = reader.ReadString('\n')
		if err != nil {
			log.Panicf("Failed to obtain user input: %w", err)
		}
		if *answer == "yes\n" || *answer == "no\n" {
			break
		}
		continue
	}
	switch {
	case *answer == "yes\n":
		_ = player.FlipCards()
		deck.DisplayCards(player.Hand)
	case *answer == "no\n":
		fmt.Println("Byyyyeeeee")
	}
}

// users to generate starting dealer, and players.
func starterSet() (deck.Dealer, []deck.Player, error) {
	splashScreen()
	dealer := deck.GenDealer()
	var players []deck.Player
	pCount := 0
	for {
		splashScreen()
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter up to 9 players.")
		sBucket, err := reader.ReadString('\n')
		if err != nil {
			return dealer, players, fmt.Errorf("Failed to read from stdin:%w", err)
		}
		pCount, err = strconv.Atoi(sBucket[:len(sBucket)-1])
		if err != nil {
			fmt.Println("Please provide a number.")
			pause(2)
			continue
		}
		if pCount <= 9 {
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
		newPlayer := deck.GenPlayer()
		players = append(players, newPlayer)
	}
	clear()
	return dealer, players, nil
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
