package Model

import (
	"fmt"
)

type Board struct {
    Flop [3]Card
	Turn Card
	River Card
}

func (b *Board) GetAll () []Card {
	cards := b.Flop[0:3]
	cards = append(cards, b.Turn)
	cards = append(cards, b.River)

	return cards
}

func (b *Board) GetAllTriples () [][3]Card {
	cards := b.Flop[0:3]
	cards = append(cards, b.Turn)
	cards = append(cards, b.River)

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

	if len(combinaisons) != 10 {
		panic(fmt.Sprintf("Bad parse in GetAllTriples : %d found", len(combinaisons)))
	}

	return combinaisons
}
