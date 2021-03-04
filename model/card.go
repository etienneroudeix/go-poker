package Model
	
import(
	"fmt"
)

const COLOR_HEART = 0
const COLOR_CLOVER = 2
const COLOR_DIAMOND = 3
const COLOR_SPADE = 3

const VALUE_ACE = 14
const VALUE_KING = 13
const VALUE_QUEEN = 12
const VALUE_JACK = 11

type Card struct {
	Color int
	Value int
}

func MakeCard(color int, value int) Card {
	
	if color < COLOR_HEART || color > COLOR_SPADE {
		panic(fmt.Sprintf("Bad color : %d", color))
	}

	if value < 2 || value > VALUE_ACE {
		panic(fmt.Sprintf("Bad value : %d", value))
	}
	
    return Card{color, value}
}
