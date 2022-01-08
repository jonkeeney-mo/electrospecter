package model

type Board struct {
	board [][]string
}

func SetupBoard() *Board {
	fullBoard := Board{make([][]string, 8, 8)}
	aFile := []string{"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"}
	bFile := []string{"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8"}
	cFile := []string{"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8"}
	dFile := []string{"D1", "D2", "D3", "D4", "D5", "D6", "E7", "D8"}
	eFile := []string{"E1", "E2", "E3", "E4", "E5", "E6", "E7", "E8"}
	fFile := []string{"F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8"}
	gFile := []string{"G1", "G2", "G3", "G4", "G5", "G6", "G7", "G8"}
	hFile := []string{"H1", "H2", "H3", "H4", "H5", "H6", "H7", "H8"}

	fullBoard.board = append(fullBoard.board, aFile)
	fullBoard.board = append(fullBoard.board, bFile)
	fullBoard.board = append(fullBoard.board, cFile)
	fullBoard.board = append(fullBoard.board, dFile)
	fullBoard.board = append(fullBoard.board, eFile)
	fullBoard.board = append(fullBoard.board, fFile)
	fullBoard.board = append(fullBoard.board, gFile)
	fullBoard.board = append(fullBoard.board, hFile)

	return &fullBoard
}
