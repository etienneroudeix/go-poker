package main

import (
	"log"
	"fmt"
	"poker/model"
)

var (
)

func main () {
	log.Println("Get improving cards")

	var board Model.Board

	board.SetFlop([3]Model.Card{
		Model.MakeCard(Model.COLOR_HEART, 10),
	   	Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
	   	Model.MakeCard(Model.COLOR_HEART, 9),
	})

	//board.SetTurn(Model.MakeCard(Model.COLOR_DIAMOND, 6))
	//board.SetRiver(Model.MakeCard(Model.COLOR_DIAMOND, 10))

	/*hands := []Model.Hand{
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
	}*/

	hand := Model.Hand{
		"4x",
		[2]Model.Card{
			Model.MakeCard(Model.COLOR_CLOVER, 8), 
			Model.MakeCard(Model.COLOR_SPADE, 2),
		},
	}

	ic := hand.GetImprovingCards(board, Model.RANK_PAIR)

	fmt.Println(len(ic), ic)
}
