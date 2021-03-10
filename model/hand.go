package Model

import(
	"fmt"
	"sort"
)

const RANK_NOTHING = 0
const RANK_PAIR = 1
const RANK_DOUBLE_PAIR = 2
const RANK_THREE_OF_KIND = 3
const RANK_STRAIGHT = 4
const RANK_FLUSH = 5
const RANK_FULL_HOUSE = 6
const RANK_FOUR_OF_A_KIND = 7
const RANK_STRAIGHT_FLUSH = 8

type Hand struct {
	Player string
    Cards [2]Card
}

type ParsedFiveCards struct {
	Rank int
	Value int
	Kickers []int
}

type ParsedHand struct {
	Hand Hand
	FiveCards ParsedFiveCards
}

func (h *Hand) Has (cardDeckValue int) bool {
	for _, card := range h.Cards {
		if card.GetDeckValue() == cardDeckValue {
			return true
		}
 	}

	return false;
}

func (h *Hand) GetBestHand (board Board) ParsedHand {		
	if board.GetState() == BOARD_STATE_PREFLOP {
		panic (fmt.Sprintf("State must not be preflop"))
	}

	hands := []ParsedFiveCards{}
	
	triples := board.GetAllTriples();
	for _, triple := range triples {
		cards := h.Cards[0:2]
		cards = append(cards, triple[0:3]...)
		var fiveCards [5]Card
		copy(fiveCards[:], cards)
		hands = append(hands, ParseFiveCards(fiveCards))
	}

	sort.SliceStable(hands, func(i, j int) bool {
		return hands[i].Compare(hands[j])
	})

	//fmt.Println(h.Player, "Best Hand : ", hands[0], "amid", )	

	return ParsedHand{*h, hands[0]}
}

func (h *Hand) GetImprovingCards (board Board, minRank int) []Card {
	cards := []Card{}
	
	currentBestHand := h.GetBestHand(board)

	fmt.Println("Current best hand", board.Flop, currentBestHand)

	for i := 0; i < 52; i++ {
		if h.Has(i) || board.Has(i) {
			continue
		}

		card := MakeCardFromDeckValue(i)

		possibleBoard := board
		possibleBoard.SetTurn(card) // todo dynamic

		possibleBestHand := h.GetBestHand(possibleBoard)
//fmt.Println("-----Card", possibleBoard.Flop, possibleBoard.Turn, possibleBestHand)
		if possibleBestHand.FiveCards.Rank < minRank {
			continue
		}

		if possibleBestHand.FiveCards.Compare(currentBestHand.FiveCards) {
			cards = append(cards, card)
		}

		// todo flush/staight draw
	}

	return cards
}

func (fiveCards *ParsedFiveCards) Compare (otherFiveCards ParsedFiveCards) bool {
	if fiveCards.Rank != otherFiveCards.Rank {
		return fiveCards.Rank > otherFiveCards.Rank
	}

	if fiveCards.Value != otherFiveCards.Value {
		return fiveCards.Value > otherFiveCards.Value
	}

	for k := range fiveCards.Kickers {
		if fiveCards.Kickers[k] == otherFiveCards.Kickers[k] {
			continue
		}

		return fiveCards.Kickers[k] > otherFiveCards.Kickers[k]
	}

	// hands are similar (color variated)
	return false
}

func ParseFiveCards (cards [5]Card) ParsedFiveCards {
	
	stack := make([]int, 13, 13)
	for _, card := range cards {
		stack[card.Value - 2]++
	}

	pairs := make([]int, 0)
	fok := 0
	tok := 0
	kickers := make([]int, 0)

	hasStraight := true
	straightTick := 0
	straightHighCard := 0

	for value, count := range stack {
		if count == 4 {
			fok = value+2;
			hasStraight = false
		}

		if count == 3 {
			tok = value+2;
			hasStraight = false
		}

		if count == 2 {
			pairs = append([]int{value+2}, pairs...)
			hasStraight = false
		}

		if count == 1 {
			kickers = append([]int{value+2}, kickers...)
			
			straightHighCard = value + 2
			straightTick++
		}

		if count == 0 {
			if straightHighCard != 0 && straightTick != 5 {
				hasStraight = false
			}
		}
	}

	flushValue := testFlush(cards)
	
	if flushValue != 0 && hasStraight {
		//fmt.Println("Has Straight Flush", flushValue, kickers)
		return ParsedFiveCards{RANK_STRAIGHT_FLUSH, flushValue, kickers}
	}
	
	if fok != 0 {
		//fmt.Println("Has Four Of a Kind", fok, kickers)
		return ParsedFiveCards{RANK_FOUR_OF_A_KIND, fok, kickers}
	}

	if tok != 0 && len(pairs) == 1 {
		//fmt.Println("Has Full House", tok, []int{pairs[0],pairs[0]})
		return ParsedFiveCards{RANK_FULL_HOUSE, tok, []int{pairs[0],pairs[0]}}
	}

	if flushValue != 0 {
		//fmt.Println("Has Flush", flushValue, kickers)
		return ParsedFiveCards{RANK_FLUSH, flushValue, kickers}
	}

	if hasStraight {
		//fmt.Println("Has Straight", straightHighCard, kickers)
		return ParsedFiveCards{RANK_STRAIGHT, straightHighCard, kickers}
	}

	if tok != 0 {
		//fmt.Println("Has Three Of a Kind", tok, kickers)
		return ParsedFiveCards{RANK_THREE_OF_KIND, tok, kickers}
	}

	if len(pairs) == 2 {
		//fmt.Println("Has Two Pairs", pairs[0], []int{pairs[1], kickers[0]})
		return ParsedFiveCards{RANK_DOUBLE_PAIR, pairs[0], []int{pairs[1], kickers[0]}}
	}

	if len(pairs) == 1 {
		//fmt.Println("Has Pair", pairs[0], kickers)
		return ParsedFiveCards{RANK_PAIR, pairs[0], kickers}
	}

	//fmt.Println("Has Nothing", kickers[0], kickers)
	return ParsedFiveCards{RANK_NOTHING, kickers[0], kickers}
}

func testFlush(cards [5]Card) (int) {
	color := cards[0].Color
	value := cards[0].Value

	for i:= 1; i<5; i++ {
		if cards[i].Color != color {
			return 0
		}

		if cards[i].Value > value {
			value = cards[i].Value
		}
	}
	
	return value
}
