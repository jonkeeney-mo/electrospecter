package model

type Piece struct {
	symbol       string
	value        int
	MoveFunction func(*BoardState) *Move
}

type PieceState struct {
	piece Piece
	color string
}

func Pawn() Piece {
	return Piece{symbol: "", value: 1, MoveFunction: MovePawn}
}

func Knight() Piece {
	return Piece{symbol: "N", value: 3, MoveFunction: MoveKnight}
}

func Bishop() Piece {
	return Piece{symbol: "B", value: 3, MoveFunction: MoveBishop}
}

func Rook() Piece {
	return Piece{symbol: "R", value: 5, MoveFunction: MoveRook}
}

func Queen() Piece {
	return Piece{symbol: "Q", value: 9, MoveFunction: MoveQueen}
}

func King() Piece {
	return Piece{symbol: "K", value: 100, MoveFunction: MoveKing}
}

func InitPieces() *BoardState {
	King := King()
	Queen := Queen()
	Rook := Rook()
	Bishop := Bishop()
	Knight := Knight()
	Pawn := Pawn()

	initBoard := make(map[string]PieceState)

	initBoard["A1"] = PieceState{piece: Rook, color: "White"}
	initBoard["B1"] = PieceState{piece: Knight, color: "White"}
	initBoard["C1"] = PieceState{piece: Bishop, color: "White"}
	initBoard["D1"] = PieceState{piece: Queen, color: "White"}
	initBoard["E1"] = PieceState{piece: King, color: "White"}
	initBoard["F1"] = PieceState{piece: Bishop, color: "White"}
	initBoard["G1"] = PieceState{piece: Knight, color: "White"}
	initBoard["H1"] = PieceState{piece: Rook, color: "White"}
	initBoard["A2"] = PieceState{piece: Pawn, color: "White"}
	initBoard["B2"] = PieceState{piece: Pawn, color: "White"}
	initBoard["C2"] = PieceState{piece: Pawn, color: "White"}
	initBoard["D2"] = PieceState{piece: Pawn, color: "White"}
	initBoard["E2"] = PieceState{piece: Pawn, color: "White"}
	initBoard["F2"] = PieceState{piece: Pawn, color: "White"}
	initBoard["G2"] = PieceState{piece: Pawn, color: "White"}
	initBoard["H2"] = PieceState{piece: Pawn, color: "White"}

	initBoard["A8"] = PieceState{piece: Rook, color: "Black"}
	initBoard["B8"] = PieceState{piece: Knight, color: "Black"}
	initBoard["C8"] = PieceState{piece: Bishop, color: "Black"}
	initBoard["D8"] = PieceState{piece: Queen, color: "Black"}
	initBoard["E8"] = PieceState{piece: King, color: "Black"}
	initBoard["F8"] = PieceState{piece: Bishop, color: "Black"}
	initBoard["G8"] = PieceState{piece: Knight, color: "Black"}
	initBoard["H8"] = PieceState{piece: Rook, color: "Black"}
	initBoard["A7"] = PieceState{piece: Pawn, color: "Black"}
	initBoard["B7"] = PieceState{piece: Pawn, color: "Black"}
	initBoard["C7"] = PieceState{piece: Pawn, color: "Black"}
	initBoard["D7"] = PieceState{piece: Pawn, color: "Black"}
	initBoard["E7"] = PieceState{piece: Pawn, color: "Black"}
	initBoard["F7"] = PieceState{piece: Pawn, color: "Black"}
	initBoard["G7"] = PieceState{piece: Pawn, color: "Black"}
	initBoard["H7"] = PieceState{piece: Pawn, color: "Black"}

	return &BoardState{State: initBoard}
}
