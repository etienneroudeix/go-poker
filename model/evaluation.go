package Model

import(
//	"fmt"
)


type Evaluation struct {
	Player string
	Hand Hand
	Outs []Card
    WinChance int
}
