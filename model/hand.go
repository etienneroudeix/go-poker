package Model

import(
	"fmt"
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

func (h *Hand) GetBestHand (board Board) {
	
	//hands := []ParsedHand{}
	
	triples := board.GetAllTriples();
	for _, triple := range triples {
		cards := h.Cards[0:2]
		cards = append(cards, triple[0:3]...)
		var fiveCards [5]Card
		copy(fiveCards[:], cards)
		ParseFiveCards(fiveCards)

		//hands = append(hands, ParsedHand{h, rank, value, kickers})
	}
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
		fmt.Println("Has Straight Flush : %d", flushValue, kickers)
		return ParsedFiveCards{RANK_STRAIGHT_FLUSH, flushValue, kickers}
	}
	
	if fok != 0 {
		fmt.Println("Has Four Of a Kind : %d", fok, kickers)
		return ParsedFiveCards{RANK_FOUR_OF_A_KIND, fok, kickers}
	}

	if tok != 0 && len(pairs) == 1 {
		fmt.Println("Has Full House : %d", tok, []int{pairs[0],pairs[0]})
		return ParsedFiveCards{RANK_FULL_HOUSE, tok, []int{pairs[0],pairs[0]}}
	}

	if flushValue != 0 {
		fmt.Println("Has Flush : %d", flushValue, kickers)
		return ParsedFiveCards{RANK_FLUSH, flushValue, kickers}
	}

	if hasStraight {
		fmt.Println("Has Straight : %d", straightHighCard, kickers)
		return ParsedFiveCards{RANK_STRAIGHT, straightHighCard, kickers}
	}

	if tok != 0 {
		fmt.Println("Has Three Of a Kind : %d", tok, kickers)
		return ParsedFiveCards{RANK_THREE_OF_KIND, tok, kickers}
	}

	if len(pairs) == 2 {
		fmt.Println("Has Two Pairs : %d", pairs[0], []int{pairs[1], kickers[0]})
		return ParsedFiveCards{RANK_THREE_OF_KIND, pairs[0], []int{pairs[1], kickers[0]}}
	}

	if len(pairs) == 1 {
		fmt.Println("Has Pair : %d", pairs[0], kickers)
		return ParsedFiveCards{RANK_THREE_OF_KIND, pairs[0], kickers}
	}

	fmt.Println("Has Nothing : %d", kickers[0], kickers)
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
