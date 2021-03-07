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

	board := Model.Board{
		[3]Model.Card{
			Model.MakeCard(Model.COLOR_HEART, 10),
			Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
			Model.MakeCard(Model.COLOR_HEART, 9),
		},
		Model.MakeCard(Model.COLOR_DIAMOND, 6),
		Model.MakeCard(Model.COLOR_DIAMOND, 10),
	}

	hands := []Model.Hand{
		{
			"4x",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_CLOVER, 10), 
				Model.MakeCard(Model.COLOR_SPADE, 10),
			},
		},
	}

	Board.ResolveRiver(board, hands)
}
