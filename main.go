package main

import (
	"fmt"
	"github.com/aybabtme/chess/board"
	"github.com/aybabtme/chess/pieces"
	"os"
)

func main() {
	b := board.NewBoard()

	printBoard := func() {
		os.Stdout.Write(b.DrawANSI())
	}

	printBoard()
	fmt.Println("starting game")
	b.Cells[0][0].Piece = pieces.Rook(board.Black)
	b.Cells[0][1].Piece = pieces.Knight(board.Black)
	b.Cells[0][2].Piece = pieces.Bishop(board.Black)
	b.Cells[0][3].Piece = pieces.Queen(board.Black)
	b.Cells[0][4].Piece = pieces.King(board.Black)
	b.Cells[0][5].Piece = pieces.Bishop(board.Black)
	b.Cells[0][6].Piece = pieces.Knight(board.Black)
	b.Cells[0][7].Piece = pieces.Rook(board.Black)
	for i := 0; i < len(b.Cells); i++ {
		b.Cells[1][i].Piece = pieces.Pawn(board.Black)
	}

	b.Cells[7][0].Piece = pieces.Rook(board.White)
	b.Cells[7][1].Piece = pieces.Knight(board.White)
	b.Cells[7][2].Piece = pieces.Bishop(board.White)
	b.Cells[7][3].Piece = pieces.Queen(board.White)
	b.Cells[7][4].Piece = pieces.King(board.White)
	b.Cells[7][5].Piece = pieces.Bishop(board.White)
	b.Cells[7][6].Piece = pieces.Knight(board.White)
	b.Cells[7][7].Piece = pieces.Rook(board.White)
	for i := 0; i < len(b.Cells); i++ {
		b.Cells[6][i].Piece = pieces.Pawn(board.White)
	}
	printBoard()
}
