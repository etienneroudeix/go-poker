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
	// todo test QF
	
	value, kickers := testFourOfAKind(cards)
	if value != 0 {
		fmt.Println("Has Four Of a Kind : %d", value)
		return ParsedFiveCards{RANK_FOUR_OF_A_KIND, value, kickers}
	}

	value, kickers = testFourOfAKind(cards)
	if value != 0 {
		fmt.Println("Has Full House : %d", value)
		return ParsedFiveCards{RANK_FULL_HOUSE, value, kickers}
	}

	// todo test F
	value, kickers = testFlush(cards)
	if value != 0 {
		fmt.Println("Has Flush : %d", value)
		return ParsedFiveCards{RANK_THREE_OF_KIND, value, kickers}
	}

	// todo test S
	value, kickers = testStraight(cards)
	if value != 0 {
		fmt.Println("Has Straight : %d", value)
		return ParsedFiveCards{RANK_STRAIGHT, value, kickers}
	}

	value, kickers = testThreeOfAKind(cards)
	if value != 0 {
		fmt.Println("Has Three Of a Kind : %d", value)
		return ParsedFiveCards{RANK_THREE_OF_KIND, value, kickers}
	}

	// todo test DP

	// todo test P

	return ParsedFiveCards{0,0,[]int{}} // TODO test nothing
}

func testFourOfAKind(cards [5]Card) (int, []int) {
	stack := make([]int, 13, 13)
	for _, card := range cards {
		stack[card.Value - 2]++
	}

	fokValue := 0
	kicker := 0

	for value, count := range stack {
		if count == 4 {
			fokValue = value
			continue
		}

		kicker = value
	}

	return fokValue, []int{kicker}
}

func testFullHouse(cards [5]Card) (int, []int) {
	stack := make([]int, 13, 13)
	for _, card := range cards {
		stack[card.Value - 2]++
	}

	tokValue := 0
	kickers := make([]int, 0)

	for value, count := range stack {
		if count == 3 {
			tokValue = value
			continue
		}

		kickers = append(kickers, value)
	}

	if tokValue == 0 {
		return 0, []int{}
	}

	if kickers[0] != kickers[1] {
		return 0, []int{}
	}

	return tokValue, kickers
}

func testFlush(cards [5]Card) (int, []int) {
	color := cards[0].Color
	value := cards[1].Value

	//fmt.Println(0, color, cards)
	for i:= 1; i<5; i++ {
		//fmt.Println(i, cards[i].Color)
		if cards[i].Color != color {
			return 0, []int{}
		}

		if cards[i].Value > value {
			value = cards[i].Value
		}
	}
	
	return value, []int{value}
}

func testStraight(cards [5]Card) (int, []int) {
	stack := make([]int, 13, 13)
	for _, card := range cards {
		stack[card.Value - 2]++
	}

	tick := 0
	highCard := 0

	for value, count := range stack {
		if count > 1 {
			return 0, []int{}
		}

		if count == 0 {
			if highCard == 0 {
				continue
			}
			return 0, []int{}
		}

		highCard = value
		tick++

		if tick == 5 {
			break
		}
	}

	return highCard, []int{}
}

func testThreeOfAKind(cards [5]Card) (int, []int) {
	stack := make([]int, 13, 13)
	for _, card := range cards {
		stack[card.Value - 2]++
	}

	tokValue := 0
	kickers := make([]int, 0)

	for value, count := range stack {
		if count == 3 {
			tokValue = value
			continue
		}

		kickers = append(kickers, value)
	}

	if tokValue == 0 {
		return 0, []int{}
	}

	if kickers[0] == kickers[1] {
		// it's a full house
		return 0, []int{}
	}

	return tokValue, kickers
}
