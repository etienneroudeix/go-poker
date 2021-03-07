package Model

import (
	"testing"
)

var (
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
