package Model

type Board struct {
    Flop [3]Card
	Turn Card
	River Card
}

func (b *Board) GetAllTriples () [][3]Card {
	cards := b.Flop[0:3]
	cards = append(cards, b.Turn)
	cards = append(cards, b.River)

	var combinaisons [][3]Card

	for i := 0; i < len(cards); i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < len(cards); k++ {
				if k == j || k == i {
					continue
				}

				combinaisons = append(combinaisons, [3]Card{
					cards[i],
					cards[j],
					cards[k],
				})
			}	
		}
	}

	return combinaisons
}
