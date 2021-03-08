package Board

import (
	"fmt"
	"sort"
	"poker/model"
)

var (
)

func EvaluateTurn (board Model.Board, hands []Model.Hand) []Model.Evaluation {
	if (board.GetState() != Model.BOARD_STATE_TURN) {
		panic (fmt.Sprintf("State must be turn, %d found", board.GetState()))
	}
	
	//Deck.CheckIntegrity(board, hands)

	bestHands := []Model.ParsedHand{}

	for _, hand := range hands {
		fmt.Printf("%s evaluation :\n", hand.Player)

		bestHands = append(bestHands, hand.GetBestHand(board))
	}

	sort.SliceStable(bestHands, func(i, j int) bool {
		return bestHands[i].FiveCards.Compare(bestHands[j].FiveCards)
	})

	evaluations := []Model.Evaluation{}

	// ...

	return evaluations
}
