package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindPairs(t *testing.T) {

	var dummyDealer Dealer

	type testCases struct {
		testSet      []Player
		handMatchSet [][]HandMatch
	}
	var tester testCases
	tester.testSet = []Player{
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
	}

	tester.handMatchSet = [][]HandMatch{
		{
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
				Name:     fmt.Sprintf("%v(%v)", fullHouse, cardIndex[2]),
				HighCard: 2,
				Value:    handValueIndex[fullHouse] + 2,
				PairType: 0,
			},
		},
		{
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
		{
			HandMatch{
				Name:     fmt.Sprintf("%v(%v)", threeOfaKind, cardIndex[3]),
				HighCard: 3,
				Value:    handValueIndex[threeOfaKind] + 3,
				PairType: 3,
			},
		},
	}

	for i, test := range tester.testSet {
		test.FindPairs()
		if !reflect.DeepEqual(test.HandMatches, tester.handMatchSet[i]) {
			t.Errorf("Test Failed:\n expected(%v)\n got     (%v)\n", tester.handMatchSet[i], test.HandMatches)
		}
	}
}
