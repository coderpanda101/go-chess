package chessboard

import (
	CON "go-chessboard/constants"
)

/*
##############################
#	A8 B8 C8 D8 E8 F8 G8 H8  #
#	A7 B7 C7 D7 E7 F7 G7 H7  #
#	A6 B6 C6 D6 E6 F6 G6 H6  #
#	A5 B5 C5 D5 E5 F5 G5 H5  #
#	A4 B4 C4 D4 E4 F4 G4 H4  #
#	A3 B3 C3 D3 E3 F3 G3 H3  #
#	A2 B2 C2 D2 E2 F2 G2 H2  #
#	A1 B1 C1 D1 E1 F1 G1 H1	 #
##############################
*/

// Piece struct represents pieces in a game of chess
type Piece struct {
	Kind      string
	Color     string
	FPosition string
	RPosition int
}

type Board struct {
	CHESSBOARD map[string]map[int]Piece
	Piece      Piece
}

// GetNewChessBoard initializes a new chessboard with the standard starting positions
func GetNewChessBoard() Board {
	board := make(map[string]map[int]Piece)

	board[CON.FILE_A] = map[int]Piece{
		1: {Kind: CON.ROOK, Color: CON.WHITEPIECE, FPosition: CON.FILE_A, RPosition: CON.RANK_1},
		2: {Kind: CON.PAWN, Color: CON.WHITEPIECE, FPosition: CON.FILE_A, RPosition: CON.RANK_2},
		3: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_A, RPosition: CON.RANK_3},
		4: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_A, RPosition: CON.RANK_4},
		5: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_A, RPosition: CON.RANK_5},
		6: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_A, RPosition: CON.RANK_6},
		7: {Kind: CON.PAWN, Color: CON.BLACKPIECE, FPosition: CON.FILE_A, RPosition: CON.RANK_7},
		8: {Kind: CON.ROOK, Color: CON.BLACKPIECE, FPosition: CON.FILE_A, RPosition: CON.RANK_8},
	}
	board[CON.FILE_B] = map[int]Piece{
		1: {Kind: CON.KNIGHT, Color: CON.WHITEPIECE, FPosition: CON.FILE_B, RPosition: CON.RANK_1},
		2: {Kind: CON.PAWN, Color: CON.WHITEPIECE, FPosition: CON.FILE_B, RPosition: CON.RANK_2},
		3: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_B, RPosition: CON.RANK_3},
		4: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_B, RPosition: CON.RANK_4},
		5: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_B, RPosition: CON.RANK_5},
		6: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_B, RPosition: CON.RANK_6},
		7: {Kind: CON.PAWN, Color: CON.BLACKPIECE, FPosition: CON.FILE_B, RPosition: CON.RANK_7},
		8: {Kind: CON.KNIGHT, Color: CON.BLACKPIECE, FPosition: CON.FILE_B, RPosition: CON.RANK_8},
	}
	board[CON.FILE_C] = map[int]Piece{
		1: {Kind: CON.BISHOP, Color: CON.WHITEPIECE, FPosition: CON.FILE_C, RPosition: CON.RANK_1},
		2: {Kind: CON.PAWN, Color: CON.WHITEPIECE, FPosition: CON.FILE_C, RPosition: CON.RANK_2},
		3: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_C, RPosition: CON.RANK_3},
		4: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_C, RPosition: CON.RANK_4},
		5: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_C, RPosition: CON.RANK_5},
		6: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_C, RPosition: CON.RANK_6},
		7: {Kind: CON.PAWN, Color: CON.BLACKPIECE, FPosition: CON.FILE_C, RPosition: CON.RANK_7},
		8: {Kind: CON.BISHOP, Color: CON.BLACKPIECE, FPosition: CON.FILE_C, RPosition: CON.RANK_8},
	}
	board[CON.FILE_D] = map[int]Piece{
		1: {Kind: CON.QUEEN, Color: CON.WHITEPIECE, FPosition: CON.FILE_D, RPosition: CON.RANK_1},
		2: {Kind: CON.PAWN, Color: CON.WHITEPIECE, FPosition: CON.FILE_D, RPosition: CON.RANK_2},
		3: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_D, RPosition: CON.RANK_3},
		4: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_D, RPosition: CON.RANK_4},
		5: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_D, RPosition: CON.RANK_5},
		6: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_D, RPosition: CON.RANK_6},
		7: {Kind: CON.PAWN, Color: CON.BLACKPIECE, FPosition: CON.FILE_D, RPosition: CON.RANK_7},
		8: {Kind: CON.QUEEN, Color: CON.BLACKPIECE, FPosition: CON.FILE_D, RPosition: CON.RANK_8},
	}
	board[CON.FILE_E] = map[int]Piece{
		1: {Kind: CON.KING, Color: CON.WHITEPIECE, FPosition: CON.FILE_E, RPosition: CON.RANK_1},
		2: {Kind: CON.PAWN, Color: CON.WHITEPIECE, FPosition: CON.FILE_E, RPosition: CON.RANK_2},
		3: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_E, RPosition: CON.RANK_3},
		4: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_E, RPosition: CON.RANK_4},
		5: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_E, RPosition: CON.RANK_5},
		6: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_E, RPosition: CON.RANK_6},
		7: {Kind: CON.PAWN, Color: CON.BLACKPIECE, FPosition: CON.FILE_E, RPosition: CON.RANK_7},
		8: {Kind: CON.KING, Color: CON.BLACKPIECE, FPosition: CON.FILE_E, RPosition: CON.RANK_8},
	}
	board[CON.FILE_F] = map[int]Piece{
		1: {Kind: CON.BISHOP, Color: CON.WHITEPIECE, FPosition: CON.FILE_F, RPosition: CON.RANK_1},
		2: {Kind: CON.PAWN, Color: CON.WHITEPIECE, FPosition: CON.FILE_F, RPosition: CON.RANK_2},
		3: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_F, RPosition: CON.RANK_3},
		4: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_F, RPosition: CON.RANK_4},
		5: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_F, RPosition: CON.RANK_5},
		6: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_F, RPosition: CON.RANK_6},
		7: {Kind: CON.PAWN, Color: CON.BLACKPIECE, FPosition: CON.FILE_F, RPosition: CON.RANK_7},
		8: {Kind: CON.BISHOP, Color: CON.BLACKPIECE, FPosition: CON.FILE_F, RPosition: CON.RANK_8},
	}
	board[CON.FILE_G] = map[int]Piece{
		1: {Kind: CON.KNIGHT, Color: CON.WHITEPIECE, FPosition: CON.FILE_G, RPosition: CON.RANK_1},
		2: {Kind: CON.PAWN, Color: CON.WHITEPIECE, FPosition: CON.FILE_G, RPosition: CON.RANK_2},
		3: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_G, RPosition: CON.RANK_3},
		4: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_G, RPosition: CON.RANK_4},
		5: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_G, RPosition: CON.RANK_5},
		6: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_G, RPosition: CON.RANK_6},
		7: {Kind: CON.PAWN, Color: CON.BLACKPIECE, FPosition: CON.FILE_G, RPosition: CON.RANK_7},
		8: {Kind: CON.KNIGHT, Color: CON.BLACKPIECE, FPosition: CON.FILE_G, RPosition: CON.RANK_8},
	}
	board[CON.FILE_H] = map[int]Piece{
		1: {Kind: CON.ROOK, Color: CON.WHITEPIECE, FPosition: CON.FILE_H, RPosition: CON.RANK_1},
		2: {Kind: CON.PAWN, Color: CON.WHITEPIECE, FPosition: CON.FILE_H, RPosition: CON.RANK_2},
		3: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_H, RPosition: CON.RANK_3},
		4: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_H, RPosition: CON.RANK_4},
		5: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_H, RPosition: CON.RANK_5},
		6: {Kind: CON.EMPTY, Color: CON.EMPTY, FPosition: CON.FILE_H, RPosition: CON.RANK_6},
		7: {Kind: CON.PAWN, Color: CON.BLACKPIECE, FPosition: CON.FILE_H, RPosition: CON.RANK_7},
		8: {Kind: CON.ROOK, Color: CON.BLACKPIECE, FPosition: CON.FILE_H, RPosition: CON.RANK_8},
	}

	return Board{CHESSBOARD: board}
}

func (b *Board) MovePiece(piece Piece) bool {

	// Check if the piece is on the board
	if _, ok := b.CHESSBOARD[piece.FPosition][piece.RPosition]; !ok {
		return false
	}

	return true
}

// Function to move Pawn
func (p *Piece) MovePawn(steps int) bool {

	// Move the pawn to the new position
	if !(((p.RPosition == CON.RANK_2) || (p.RPosition == CON.RANK_7)) && steps == 2) {
		return false // Cannot move this Pawn by 2 steps
	}

	if steps > 2 || steps < 0 {
		return false // Invalid steps
	}

	p.RPosition += steps // moving the pawn by specified steps
	return true
}

func (p *Piece) MoveRook(direction string, steps int) bool {
	return true
}
