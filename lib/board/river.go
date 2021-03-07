package Board

import (
	"fmt"
	"sort"
	"poker/model"
	"poker/lib/deck"
)

var (
)

func ResolveRiver (board Model.Board, hands []Model.Hand) Model.ParsedHand {
	Deck.CheckIntegrity(board, hands)
	
	bestHands := []Model.ParsedHand{}

	for _, hand := range hands {
		fmt.Printf("%s best hand : \n", hand.Player)

		bestHands = append(bestHands, hand.GetBestHand(board))
	}

	sort.SliceStable(bestHands, func(i, j int) bool {
		return bestHands[i].FiveCards.Compare(bestHands[j].FiveCards)
	})

	return bestHands[0]
}
