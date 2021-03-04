package main

import (
	"log"
	"poker/lib/board"
	"poker/model"
)

var (
)

func main () {
	log.Println("Welcome")

	board := Model.Board{
		[3]Model.Card{
			Model.MakeCard(Model.COLOR_HEART, 10),
			Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
			Model.MakeCard(Model.COLOR_HEART, 9),
		},
		Model.MakeCard(Model.COLOR_DIAMOND, 10),
		Model.MakeCard(Model.COLOR_DIAMOND, 6),
	}

	hands := []Model.Hand{
		{
			"QF",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_HEART, Model.VALUE_QUEEN), 
				Model.MakeCard(Model.COLOR_HEART, 8),
			},
		},
		{
			"4x",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_CLOVER, 10), 
				Model.MakeCard(Model.COLOR_SPADE, 10),
			},
		},
		{
			"Full",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_CLOVER, 6), 
				Model.MakeCard(Model.COLOR_SPADE, 6),
			},
		},
		{
			"Color",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_HEART, Model.VALUE_ACE), 
				Model.MakeCard(Model.COLOR_HEART, 3),
			},
		},
		{
			"Straight",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_SPADE, Model.VALUE_QUEEN), 
				Model.MakeCard(Model.COLOR_DIAMOND, Model.VALUE_KING),
			},
		},
		{
			"3x",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_SPADE, 10), 
				Model.MakeCard(Model.COLOR_DIAMOND, 3),
			},
		},
		{
			"2x2",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_SPADE, 7), 
				Model.MakeCard(Model.COLOR_DIAMOND, 7),
			},
		},
		{
			"2x",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_SPADE, 7), 
				Model.MakeCard(Model.COLOR_DIAMOND, 5),
			},
		},
	}

	Board.ResolveRiver(board, hands)
}
