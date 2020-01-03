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
	}

	for i, test := range tester.testSet {
		test.FindPairs()
		if !reflect.DeepEqual(test.HandMatches, tester.handMatchSet[i]) {
			t.Errorf("Test Failed:\n expected(%v)\n got     (%v)\n", tester.handMatchSet[i], test.HandMatches)
		}
	}
}
