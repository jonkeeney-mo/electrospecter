package main

import (
	"electrospecter/model"
	"fmt"
)

func main() {

	state := model.InitPieces()
	board := model.SetupBoard()

	fmt.Println(state)
	fmt.Println(board)
}
