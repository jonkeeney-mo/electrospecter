package model

// BoardState is really the backbone of state for the game. Board, meanwhile, is just a chessboard.
type BoardState struct {
	State map[string]PieceState

	// would love a better way to indicate black/white but w/e
	WhoseMove string
	Check     bool

	// TODO: pawn move events will trigger update to board status to allow EP
	// system will then look for pawns on advanced rank next to opponent pawns
	EnPassant bool

	// castling flags also triggered by move events
	WhiteCanCastle bool
	BlackCanCastle bool
}
