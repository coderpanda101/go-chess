package main

import (
	"go-chessboard/chessboard"
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

func main() {

	board := make(map[string]map[int]chessboard.Piece)

	board["a"] = map[int]chessboard.Piece{
		1: {Kind: "rook", Color: "white"},
		2: {Kind: "pawn", Color: "white"},
		3: {Kind: "", Color: ""},
		4: {Kind: "", Color: ""},
		5: {Kind: "", Color: ""},
		6: {Kind: "", Color: ""},
		7: {Kind: "pawn", Color: "black"},
		8: {Kind: "rook", Color: "black"},
	}
	board["b"] = map[int]chessboard.Piece{
		1: {Kind: "knight", Color: "white"},
		2: {Kind: "pawn", Color: "white"},
		3: {Kind: "", Color: ""},
		4: {Kind: "", Color: ""},
		5: {Kind: "", Color: ""},
		6: {Kind: "", Color: ""},
		7: {Kind: "pawn", Color: "black"},
		8: {Kind: "knight", Color: "black"},
	}
	board["c"] = map[int]chessboard.Piece{
		1: {Kind: "bishop", Color: "white"},
		2: {Kind: "pawn", Color: "white"},
		3: {Kind: "", Color: ""},
		4: {Kind: "", Color: ""},
		5: {Kind: "", Color: ""},
		6: {Kind: "", Color: ""},
		7: {Kind: "pawn", Color: "black"},
		8: {Kind: "bishop", Color: "black"},
	}
	board["d"] = map[int]chessboard.Piece{
		1: {Kind: "queen", Color: "white"},
		2: {Kind: "pawn", Color: "white"},
		3: {Kind: "", Color: ""},
		4: {Kind: "", Color: ""},
		5: {Kind: "", Color: ""},
		6: {Kind: "", Color: ""},
		7: {Kind: "pawn", Color: "black"},
		8: {Kind: "queen", Color: "black"},
	}
	board["e"] = map[int]chessboard.Piece{
		1: {Kind: "king", Color: "white"},
		2: {Kind: "pawn", Color: "white"},
		3: {Kind: "", Color: ""},
		4: {Kind: "", Color: ""},
		5: {Kind: "", Color: ""},
		6: {Kind: "", Color: ""},
		7: {Kind: "pawn", Color: "black"},
		8: {Kind: "king", Color: "black"},
	}
	board["f"] = board["c"]
	board["g"] = board["b"]
	board["h"] = board["a"]

}
