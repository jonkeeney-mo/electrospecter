package model

type Move struct {
	MovingPiece   Piece
	FileChange    int
	RankChange    int
	CapturedPiece Piece
}
