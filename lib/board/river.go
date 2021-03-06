package Board

import (
	"fmt"
	"sort"
	"poker/model"
)

var (
)

func ResolveRiver (board Model.Board, hands []Model.Hand) {
	fmt.Println("Resolving")

	bestHands := []Model.ParsedHand{}

	for _, hand := range hands {
		fmt.Printf("%s best hand : \n", hand.Player)

		bestHands = append(bestHands, hand.GetBestHand(board))
	}

	sort.SliceStable(bestHands, func(i, j int) bool {
		return bestHands[i].FiveCards.Compare(bestHands[j].FiveCards)
	})

	fmt.Println("Winner : ", bestHands[0].Hand.Player)
}
