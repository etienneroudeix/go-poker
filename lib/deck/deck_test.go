package Deck

import (
	"testing"
	"poker/model"
)

var (
)

func TestCheckIntegrity(t *testing.T) {
	var board Model.Board

	board.SetFlop([3]Model.Card{
		Model.MakeCard(Model.COLOR_HEART, 10),
	   	Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
	   	Model.MakeCard(Model.COLOR_HEART, 9),
	})

	board.SetTurn(Model.MakeCard(Model.COLOR_DIAMOND, 6))
	board.SetRiver(Model.MakeCard(Model.COLOR_DIAMOND, 5))

	hands := []Model.Hand{
		{
			"QF",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_HEART, Model.VALUE_QUEEN), 
				Model.MakeCard(Model.COLOR_HEART, 8),
			},
		},
	}

    CheckIntegrity(board, hands)	
}

func TestCheckIntegrityFailure(t *testing.T) {
	var board Model.Board

	board.SetFlop([3]Model.Card{
		Model.MakeCard(Model.COLOR_HEART, 10),
	   	Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
	   	Model.MakeCard(Model.COLOR_HEART, 9),
	})

	board.SetTurn(Model.MakeCard(Model.COLOR_DIAMOND, 6))
	board.SetRiver(Model.MakeCard(Model.COLOR_DIAMOND, 5))

	hands := []Model.Hand{
		{
			"QF",
			[2]Model.Card{
				Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK), 
				Model.MakeCard(Model.COLOR_HEART, 8),
			},
		},
	}

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()

    CheckIntegrity(board, hands)

	
}