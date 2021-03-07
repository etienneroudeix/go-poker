package Deck

import (
	"fmt"
	"poker/model"
)

var (
)

func CheckIntegrity(board Model.Board, hands []Model.Hand) {
	fmt.Println("Checking integrity")

	deck := make([]int, 52, 52)

	cards := board.GetAll()
	for _, hand := range hands {
		cards = append(cards, hand.Cards[0:2]...)
	}

	for _, card := range cards {
		if deck[card.GetDeckValue()] > 0 {
			panic(fmt.Sprintf("Deck integrity compromised : %d", card.GetDeckValue()))
		}
		deck[card.GetDeckValue()]++
	}
}