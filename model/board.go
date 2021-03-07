package Model

import (
	"fmt"
)

const BOARD_STATE_PREFLOP = 0
const BOARD_STATE_FLOP = 1
const BOARD_STATE_TURN = 2
const BOARD_STATE_RIVER = 3

type Board struct {
    Flop *[3]Card
	Turn *Card
	River *Card
}

func (b *Board) GetState () int {
	if b.River != nil {
		return BOARD_STATE_RIVER
	}

	if b.Turn != nil {
		return BOARD_STATE_TURN
	}

	if b.Flop != nil {
		return BOARD_STATE_FLOP
	}

	return BOARD_STATE_PREFLOP
}

func (b *Board) SetFlop (flop [3]Card) {
	if (b.GetState() != BOARD_STATE_PREFLOP) {
		panic (fmt.Sprintf("Flop cannot be set from state %d", b.GetState()))
	}

	b.Flop = &flop
}

func (b *Board) SetTurn (turn Card) {
	if (b.GetState() != BOARD_STATE_FLOP) {
		panic (fmt.Sprintf("Turn cannot be set from state %d", b.GetState()))
	}

	b.Turn = &turn
}

func (b *Board) SetRiver (river Card) {
	if (b.GetState() != BOARD_STATE_TURN) {
		panic (fmt.Sprintf("River cannot be set from state %d", b.GetState()))
	}

	b.River = &river
}

func (b *Board) GetAll () []Card {
	cards := []Card{}

	if b.GetState() >= BOARD_STATE_FLOP {
		cards = append(cards, []Card{b.Flop[0], b.Flop[1], b.Flop[2]}...)
	}

	if b.GetState() >= BOARD_STATE_TURN {
		cards = append(cards, *b.Turn)
	}

	if b.GetState() == BOARD_STATE_RIVER {
		cards = append(cards, *b.River)
	}

	return cards
}

func (b *Board) GetAllTriples () [][3]Card {
	cards := b.GetAll()

	expectedTriplesCount := 10;
	switch b.GetState() {
	case BOARD_STATE_PREFLOP:
		expectedTriplesCount = 0
	case BOARD_STATE_FLOP:
		expectedTriplesCount = 1
	case BOARD_STATE_TURN:
		expectedTriplesCount = 4
	}

	var combinaisons [][3]Card

	for i := 0; i < len(cards); i++ {
		for j := 0; j < i; j++ {
			for k := i+1; k < len(cards); k++ {
				combinaisons = append(combinaisons, [3]Card{
					cards[i],
					cards[j],
					cards[k],
				})
			}	
		}
	}

	if len(combinaisons) != expectedTriplesCount {
		panic(fmt.Sprintf("Bad parse in GetAllTriples : %d found", len(combinaisons)))
	}

	return combinaisons
}

