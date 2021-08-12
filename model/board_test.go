package Model

import (
	"testing"
)

var (
	test = true
)

func TestGetAllTriples(t *testing.T) {
	var board Board

	if len(board.GetAllTriples()) != 0 {
		t.Errorf("Expected no result")
	}

	board.SetFlop([3]Card{
		MakeCard(COLOR_HEART, 10),
	   	MakeCard(COLOR_HEART, VALUE_JACK),
	   	MakeCard(COLOR_HEART, 9),
	})

	if len(board.GetAllTriples()) != 1 {
		t.Errorf("Expected 1 result")
	}

	board.SetTurn(MakeCard(COLOR_DIAMOND, 6))

	if len(board.GetAllTriples()) != 4 {
		t.Errorf("Expected 4 results")
	}

	board.SetRiver(MakeCard(COLOR_DIAMOND, 10))

	if len(board.GetAllTriples()) != 10 {
		t.Errorf("Expected 10 result")
	}
}

func TestGetWinner (t *testing.T) {
	var board Board

	board.SetFlop([3]Card{
		MakeCard(COLOR_HEART, 10),
	   	MakeCard(COLOR_HEART, VALUE_JACK),
	   	MakeCard(COLOR_HEART, 9),
	})

	board.SetTurn(MakeCard(COLOR_DIAMOND, 6))
	board.SetRiver(MakeCard(COLOR_DIAMOND, 10))

	hands := []Hand{
		{
			"QF",
			[2]Card{
				MakeCard(COLOR_HEART, VALUE_QUEEN), 
				MakeCard(COLOR_HEART, 8),
			},
		},
		{
			"4x",
			[2]Card{
				MakeCard(COLOR_CLOVER, 10), 
				MakeCard(COLOR_SPADE, 10),
			},
		},
		{
			"Full",
			[2]Card{
				MakeCard(COLOR_CLOVER, 6), 
				MakeCard(COLOR_SPADE, 6),
			},
		},
		{
			"Color",
			[2]Card{
				MakeCard(COLOR_HEART, VALUE_ACE), 
				MakeCard(COLOR_HEART, 3),
			},
		},
		{
			"Straight",
			[2]Card{
				MakeCard(COLOR_SPADE, VALUE_QUEEN), 
				MakeCard(COLOR_DIAMOND, VALUE_KING),
			},
		},
		/*{
			"3x",
			[2]Card{
				MakeCard(COLOR_SPADE, 10), 
				MakeCard(COLOR_DIAMOND, 3),
			},
		},*/
		{
			"2x2",
			[2]Card{
				MakeCard(COLOR_SPADE, 7), 
				MakeCard(COLOR_DIAMOND, 7),
			},
		},
		{
			"2x",
			[2]Card{
				MakeCard(COLOR_CLOVER, 7), 
				MakeCard(COLOR_DIAMOND, 5),
			},
		},
	}

	winner := board.GetWinner(hands)

	if winner.Hand.Player != "QF" {
		t.Errorf("Bad winner")
	}

	if winner.FiveCards.Rank != RANK_STRAIGHT_FLUSH {
		t.Errorf("Bad rank")
	}
}

func TestCheckIntegrity(t *testing.T) {
	var board Board

	board.SetFlop([3]Card{
		MakeCard(COLOR_HEART, 10),
	   	MakeCard(COLOR_HEART, VALUE_JACK),
	   	MakeCard(COLOR_HEART, 9),
	})

	board.SetTurn(MakeCard(COLOR_DIAMOND, 6))
	board.SetRiver(MakeCard(COLOR_DIAMOND, 5))

	hands := []Hand{
		{
			"QF",
			[2]Card{
				MakeCard(COLOR_HEART, VALUE_QUEEN), 
				MakeCard(COLOR_HEART, 8),
			},
		},
	}

    board.CheckIntegrity(hands)	
}

func TestCheckIntegrityFailure(t *testing.T) {
	var board Board

	board.SetFlop([3]Card{
		MakeCard(COLOR_HEART, 10),
	   	MakeCard(COLOR_HEART, VALUE_JACK),
	   	MakeCard(COLOR_HEART, 9),
	})

	board.SetTurn(MakeCard(COLOR_DIAMOND, 6))
	board.SetRiver(MakeCard(COLOR_DIAMOND, 5))

	hands := []Hand{
		{
			"QF",
			[2]Card{
				MakeCard(COLOR_HEART, VALUE_JACK), 
				MakeCard(COLOR_HEART, 8),
			},
		},
	}

	defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()

    board.CheckIntegrity(hands)

	
}
