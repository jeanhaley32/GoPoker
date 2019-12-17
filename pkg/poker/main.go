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
	dealer, _ := starterSet()
	var players []*deck.Player

	players = append(players, dealer.Players...)
	for _, player := range players {
		fmt.Println(player.Name, len(player.Hand.Cards))
		player.Hand.FlipCards()
		player.Hand.DisplayCards()
	}
}

// users to generate starting dealer, and players.
func starterSet() (*deck.Dealer, error) {
	splashScreen()
	var dealer deck.Dealer
	dealer.GenDeck()
	pCount := 0
	for {
		splashScreen()
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter up to 9 players.")
		sBucket, err := reader.ReadString('\n')
		if err != nil {
			return &dealer, fmt.Errorf("Failed to read from stdin:%w", err)
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
		_ = dealer.GenPlayer()
	}
	clear()
	return &dealer, nil
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
