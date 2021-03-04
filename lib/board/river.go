package Board

import (
	"fmt"
	"poker/model"
)

var (
)

func ResolveRiver (board Model.Board, hands []Model.Hand) {
	fmt.Println("Resolving")

	for _, hand := range hands {
		fmt.Printf("%s best hand : \n", hand.Player)

		hand.GetBestHand(board)
	}

	fmt.Printf("Winner : %s", "Bob")
}
