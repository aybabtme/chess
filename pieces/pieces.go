package pieces

import (
	"github.com/aybabtme/chess/board"
)

type Piece struct {
	color board.Color
	char  rune
}

func (p Piece) Color() board.Color {
	return p.color
}

func (p Piece) Char() rune {
	return p.char
}

func King(color board.Color) Piece {
	p := Piece{color: color}
	switch color {
	case board.Black:
		p.char = blackKing
	case board.White:
		p.char = whiteKing
	}
	return p
}

func Queen(color board.Color) Piece {
	p := Piece{color: color}
	switch color {
	case board.Black:
		p.char = blackQueen
	case board.White:
		p.char = whiteQueen
	}
	return p
}

func Rook(color board.Color) Piece {
	p := Piece{color: color}
	switch color {
	case board.Black:
		p.char = blackRook
	case board.White:
		p.char = whiteRook
	}
	return p
}

func Bishop(color board.Color) Piece {
	p := Piece{color: color}
	switch color {
	case board.Black:
		p.char = blackBishop
	case board.White:
		p.char = whiteBishop
	}
	return p
}

func Knight(color board.Color) Piece {
	p := Piece{color: color}
	switch color {
	case board.Black:
		p.char = blackKnight
	case board.White:
		p.char = whiteKnight
	}
	return p
}

func Pawn(color board.Color) Piece {
	p := Piece{color: color}
	switch color {
	case board.Black:
		p.char = blackPawn
	case board.White:
		p.char = whitePawn
	}
	return p
}

const (
	whiteKing   = '♔'
	whiteQueen  = '♕'
	whiteRook   = '♖'
	whiteBishop = '♗'
	whiteKnight = '♘'
	whitePawn   = '♙'
	blackKing   = '♚'
	blackQueen  = '♛'
	blackRook   = '♜'
	blackBishop = '♝'
	blackKnight = '♞'
	blackPawn   = '♟'
)
