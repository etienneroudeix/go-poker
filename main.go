package main

import (
	"log"
	"poker/lib/board"
	"poker/model"
)

var (
)

func main () {
	log.Println("Test Draw")

	var board Model.Board

	board.SetFlop([3]Model.Card{
		Model.MakeCard(Model.COLOR_HEART, 10),
	   	Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
	   	Model.MakeCard(Model.COLOR_HEART, 9),
	})

	board.SetTurn(Model.MakeCard(Model.COLOR_DIAMOND, 6))
	//board.SetRiver(Model.MakeCard(Model.COLOR_DIAMOND, 10))

	hands := []Model.Hand{
		{
			"4x",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_CLOVER, 10), 
				Model.MakeCard(Model.COLOR_SPADE, 10),
			},
		},
		{
			"4x",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_CLOVER, 4), 
				Model.MakeCard(Model.COLOR_SPADE, 3),
			},
		},
	}

	Board.EvaluateTurn(board, hands)
}
