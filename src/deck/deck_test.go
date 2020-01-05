package deck

import (
	"fmt"
	"testing"
)

func TestFindPairs(t *testing.T) {

	var dummyDealer Dealer

	type testCase struct {
		testPlayer   Player
		handMatchSet []HandMatch
	}

	cases := []testCase{
		testCase{
			Player{
				Dealer: &dummyDealer,
				Hand: CardCollection{
					Cards: []Card{
						Card{Rank: 2},
						Card{Rank: 2},
						Card{Rank: 2},
						Card{Rank: 2},
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 3},
					},
				},
			},
			[]HandMatch{
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", fourOfaKind, cardIndex[2]),
					HighCard: 2,
					Value:    handValueIndex[fourOfaKind] + 2,
					PairType: 4,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[threeOfaKind] + 3,
					PairType: 3,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", fullHouse, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[fullHouse] + 3,
					PairType: 0,
				},
			},
		},
		testCase{
			Player{
				Dealer: &dummyDealer,
				Hand: CardCollection{
					Cards: []Card{
						Card{Rank: 2},
						Card{Rank: 2},
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 4},
						Card{Rank: 4},
						Card{Rank: 4},
					},
				},
			},
			[]HandMatch{
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", twoPair, cardIndex[2]),
					HighCard: 2,
					Value:    handValueIndex[twoPair] + 2,
					PairType: 2,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", twoPair, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[twoPair] + 3,
					PairType: 2,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[4]),
					HighCard: 4,
					Value:    handValueIndex[threeOfaKind] + 4,
					PairType: 3,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", fullHouse, cardIndex[4]),
					HighCard: 4,
					Value:    handValueIndex[fullHouse] + 4,
					PairType: 0,
				},
			},
		},
		testCase{
			Player{
				Dealer: &dummyDealer,
				Hand: CardCollection{
					Cards: []Card{
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 14},
						Card{Rank: 14},
						Card{Rank: 12},
						Card{Rank: 4},
					},
				},
			},
			[]HandMatch{
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", fullHouse, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[fullHouse] + 3,
					PairType: 0,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[threeOfaKind] + 3,
					PairType: 3,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", twoPair, cardIndex[14]),
					HighCard: 14,
					Value:    handValueIndex[twoPair] + 14,
					PairType: 2,
				},
			},
		},
		testCase{
			Player{
				Dealer: &dummyDealer,
				Hand: CardCollection{
					Cards: []Card{
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 14},
						Card{Rank: 13},
						Card{Rank: 6},
						Card{Rank: 2},
					},
				},
			},
			[]HandMatch{
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[threeOfaKind] + 3,
					PairType: 3,
				},
			},
		},
		testCase{
			Player{
				Dealer: &dummyDealer,
				Hand: CardCollection{
					Cards: []Card{
						Card{Rank: 9},
						Card{Rank: 9},
						Card{Rank: 13},
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 3},
					},
				},
			},
			[]HandMatch{
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", fourOfaKind, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[fourOfaKind] + 3,
					PairType: 4,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", twoPair, cardIndex[9]),
					HighCard: 9,
					Value:    handValueIndex[twoPair] + 9,
					PairType: 2,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", fullHouse, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[fullHouse] + 3,
					PairType: 0,
				},
			},
		},
		testCase{
			Player{
				Dealer: &dummyDealer,
				Hand: CardCollection{
					Cards: []Card{
						Card{Rank: 5},
						Card{Rank: 5},
						Card{Rank: 5},
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 3},
						Card{Rank: 2},
					},
				},
			},
			[]HandMatch{
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[5]),
					HighCard: 5,
					Value:    handValueIndex[threeOfaKind] + 5,
					PairType: 3,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[3]),
					HighCard: 3,
					Value:    handValueIndex[threeOfaKind] + 3,
					PairType: 3,
				},
				HandMatch{
					Name:     fmt.Sprintf("%v(%v)", fullHouse, cardIndex[5]),
					HighCard: 5,
					Value:    handValueIndex[fullHouse] + 5,
					PairType: 0,
				},
			},
		},
	}

	fmt.Println("Hello roger")
	for _, c := range cases {
		c.testPlayer.FindPairs()
		if !compare(c.testPlayer.HandMatches, c.handMatchSet) {
			t.Errorf("Test Failed:\n expected(%v)\n got     (%v)\n", c.handMatchSet, c.testPlayer.HandMatches)
		}
	}
}

// compare compares two un-ordered slices to check if they contain the same elements, returns a boolean.

func compare(a, b []HandMatch) bool {
	if len(a) != len(b) {
		return false
	}
	matches := 0
	for _, matchA := range a {
		for _, matchB := range b {
			if matchA == matchB {
				matches++
			}
		}
	}
	if matches != len(a) {
		return false
	}
	return true
}
