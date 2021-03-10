package Model

import (
	"fmt"
	"sort"
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

func (b *Board) Has (cardDeckValue int) bool {
	for _, card := range b.Flop {
		if card.GetDeckValue() == cardDeckValue {
			return true
		}
	 }
	 
	 if (b.Turn != nil && b.Turn.GetDeckValue() == cardDeckValue) || (b.River != nil && b.River.GetDeckValue() == cardDeckValue) {
		 return true
	 }

	return false;
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

func (board *Board) CheckIntegrity(hands []Hand) {
	fmt.Println("Checking integrity")

	deck := make([]int, 52, 52)

	cards := board.GetAll()
	for _, hand := range hands {
		cards = append(cards, hand.Cards[0:2]...)
	}

	for _, card := range cards {
		if deck[card.GetDeckValue()] > 0 {
			panic(fmt.Sprintf("Deck integrity compromised : %d", card.GetDeckValue()))
		}
		deck[card.GetDeckValue()]++
	}
}

/*func (board *Board) EvaluateTurn (hands []Hand) []Evaluation {
	if (board.GetState() != BOARD_STATE_TURN) {
		panic (fmt.Sprintf("State must be turn, %d found", board.GetState()))
	}
	
	board.CheckIntegrity(hands)

	bestHands := []ParsedHand{}

	for i := 0; i < 52; i++ {
		
	}

	cards := board.GetAll()
	for i, card := range deck {
		cards = append(cards, hand.Cards[0:2]...)
	}

	for _, hand := range hands {
		fmt.Printf("%s get outs :\n", hand.Player)

		bestHands = append(bestHands, hand.GetBestHand(*board))			
	}

	sort.SliceStable(bestHands, func(i, j int) bool {
		return bestHands[i].FiveCards.Compare(bestHands[j].FiveCards)
	})

	evaluations := []Evaluation{}

	// ...

	return evaluations
}*/

func (board *Board) GetWinner (hands []Hand) ParsedHand {
	if board.GetState() != BOARD_STATE_RIVER {
		panic (fmt.Sprintf("State must be river, %d found", board.GetState()))
	}

	board.CheckIntegrity(hands)

	bestHands := []ParsedHand{}

	for _, hand := range hands {
		fmt.Printf("%s best hand : \n", hand.Player)

		bestHands = append(bestHands, hand.GetBestHand(*board))
	}

	sort.SliceStable(bestHands, func(i, j int) bool {
		return bestHands[i].FiveCards.Compare(bestHands[j].FiveCards)
	})

	return bestHands[0]
}
