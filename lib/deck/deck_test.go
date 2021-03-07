package Deck

import (
	"testing"
	"poker/model"
)

var (
)

func TestCheckIntegrity(t *testing.T) {
	board := Model.Board{
		[3]Model.Card{
			Model.MakeCard(Model.COLOR_HEART, 10),
			Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
			Model.MakeCard(Model.COLOR_HEART, 9),
		},
		Model.MakeCard(Model.COLOR_DIAMOND, 5),
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
	}

    CheckIntegrity(board, hands)	
}

func TestCheckIntegrityFailure(t *testing.T) {
	board := Model.Board{
		[3]Model.Card{
			Model.MakeCard(Model.COLOR_HEART, 10),
			Model.MakeCard(Model.COLOR_HEART, Model.VALUE_JACK),
			Model.MakeCard(Model.COLOR_HEART, 9),
		},
		Model.MakeCard(Model.COLOR_DIAMOND, 5),
		Model.MakeCard(Model.COLOR_DIAMOND, 6),
	}

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