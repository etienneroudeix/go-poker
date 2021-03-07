package Board

import (
	"testing"
	"poker/model"
)

func TestResolveRiver (t *testing.T) {
	var board Model.Board

	board.SetFlop([3]Model.Card{
		Model.MakeCard(Model.COLOR_HEART, 10),
	   	Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
	   	Model.MakeCard(Model.COLOR_HEART, 9),
	})

	board.SetTurn(Model.MakeCard(Model.COLOR_DIAMOND, 6))
	board.SetRiver(Model.MakeCard(Model.COLOR_DIAMOND, 10))

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
		/*{
			"3x",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_SPADE, 10), 
				Model.MakeCard(Model.COLOR_DIAMOND, 3),
			},
		},*/
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
				Model.MakeCard(Model.COLOR_CLOVER, 7), 
				Model.MakeCard(Model.COLOR_DIAMOND, 5),
			},
		},
	}

	winner := ResolveRiver(board, hands)

	if winner.Hand.Player != "QF" {
		t.Errorf("Bad winner")
	}

	if winner.FiveCards.Rank != Model.RANK_STRAIGHT_FLUSH {
		t.Errorf("Bad rank")
	}
}
