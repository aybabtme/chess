package board

import (
	"bytes"
	"fmt"
	"github.com/aybabtme/rgbterm"
	"log"
)

type Board struct {
	Cells [8][8]Cell
}

func NewBoard() *Board {
	b := Board{}
	for i, rows := range b.Cells {
		for j := range rows {
			c := newCell(i, j)
			b.Cells[i][j] = c
		}
	}
	return &b
}

func (b Board) DrawANSI() []byte {
	buf := bytes.NewBuffer(nil)
	fmt.Fprint(buf, "_")
	for i := 0; i < len(b.Cells); i++ {
		buf.WriteRune(' ')
		buf.Write(itoa(i + 1))
	}
	for i, rows := range b.Cells {
		fmt.Fprintf(buf, "\n%d", 8-i)
		for _, cell := range rows {
			buf.Write(cell.DrawANSI())
		}
		fmt.Fprintf(buf, "%d", 8-i)
	}
	fmt.Fprint(buf, "\n ")
	for i := 0; i < len(b.Cells); i++ {
		buf.WriteRune(' ')
		buf.Write(itoa(i + 1))
	}
	buf.WriteRune('\n')
	return buf.Bytes()
}

type Cell struct {
	ID       int
	Row, Col int
	Color    Color
	Piece    interface {
		Char() rune
		Color() Color
	}
}

func newCell(i, j int) Cell {
	c := Cell{
		ID:  (i * 8) + j,
		Row: i, Col: j,
	}
	if (i+j)%2 == 0 {
		c.Color = White
	} else {
		c.Color = Black
	}
	return c
}

func (c Cell) String() string {
	a, b := xyToAlg(c.Col, c.Row)
	return string(a) + fmt.Sprintf("%d (%d-%s)", b, c.ID, c.Color)
}

func (c Cell) DrawANSI() []byte {
	var piece []byte
	var fr, fg, fb uint8
	if c.Piece != nil {
		char := string(c.Piece.Char())
		if len([]rune(char)) < 2 {
			char = " " + char
		} else if len([]rune(char)) > 2 {
			char = char[:2]
		}
		piece = []byte(char)
		fr, fg, fb = c.Piece.Color().highRGB()
	} else {
		piece = []byte("  ")
	}
	br, bg, bb := c.Color.rgb()
	return rgbterm.Bytes(piece, fr, fg, fb, br, bg, bb)
}

type Color int

const (
	invalid = iota
	Black
	White
)

func (c Color) String() string {
	switch c {
	case Black:
		return "black"
	case White:
		return "white"
	}
	return fmt.Sprintf("invalid color (%d)", c)
}

func (c Color) rgb() (r, g, b uint8) {
	switch c {
	case Black:
		return 96, 96, 96
	case White:
		return 190, 190, 190
	}
	panic(c.String())
}

func (c Color) highRGB() (r, g, b uint8) {
	switch c {
	case Black:
		return 0, 0, 0
	case White:
		return 127, 127, 220
	}
	panic(c.String())
}

func atoi(a byte) int {
	if a < 'a' || a > 'h' {
		log.Panicf("%q is out of the [a-z] range", a)
	}
	return int(a-'a') + 1
}

func itoa(i int) []byte {
	if i < 1 || i > 8 {
		log.Panicf("%d is out of the [1-8] range", i)
	}
	return []byte{byte(i-1) + 'a'}
}

func algToXY(a rune, b int) (x, y int) {
	return atoi(byte(a)), 8 - b - 1
}

func xyToAlg(x, y int) (a rune, b int) {
	return rune(x) + 'a', 8 - y
}
