package game

import (
	"math/rand"
)

// Piece is a constant for a shape of piece. There are 7 classic pieces like L, and O
type Shape int

// Various values that the pieces can be
const (
	IPiece Shape = iota
	JPiece
	LPiece
	OPiece
	SPiece
	TPiece
	ZPiece
)

type Coords struct {
	R, C int
}

type Piece struct {
	Shape
	Block
	Points []Coords
}

func (p *Piece) copy() *Piece {
	np := *p
	copy(np.Points, p.Points)
	return &np
}

func (p *Piece) rotate() *Piece {
	pivot := p.Points[1]
	np := p.copy()
	copy(np.Points, p.Points)
	for i, v := range p.Points {
		if i == 1 {
			continue
		}
		rowDiff := pivot.R - v.R
		colDiff := pivot.C - v.C
		// New Row is how far the old column was away from the pivot but multiplied by -1
		// New Column is how far away the old row was from the pivot
		np.Points[i] = Coords{pivot.R + (colDiff * -1), pivot.C + rowDiff}
	}
	return np
}

func GenerateRandomPiece() *Piece {
	i := rand.Intn(7)
	switch i {
	case 0:
		return NewLPiece()
	case 1:
		return NewIPiece()
	case 2:
		return NewOPiece()
	case 3:
		return NewTPiece()
	case 4:
		return NewSPiece()
	case 5:
		return NewZPiece()
	default:
		return NewJPiece()
	}
}

func (p *Piece) move(r, c int) *Piece {
	np := p.copy()
	for i, v := range p.Points {
		np.Points[i].R = v.R + r
		np.Points[i].C = v.C + c
	}
	return np
}

func (p *Piece) moveDown() *Piece {
	return p.move(1, 0)
}
func (p *Piece) moveRight() *Piece {
	return p.move(0, 1)
}
func (p *Piece) moveLeft() *Piece {
	return p.move(0, -1)
}

func NewLPiece() *Piece {
	p := &Piece{
		Shape: LPiece,
		Block: LightBlue,
		Points: []Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{0, 0},
		},
	}
	return p
}

func NewIPiece() *Piece {
	p := &Piece{
		Shape: IPiece,
		Block: Blue,
		Points: []Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{1, 3},
		},
	}
	return p
}

func NewOPiece() *Piece {
	p := &Piece{
		Shape: OPiece,
		Block: Pink,
		Points: []Coords{
			{1, 0},
			{1, 1},
			{0, 0},
			{0, 1},
		},
	}
	return p
}

func NewTPiece() *Piece {
	p := &Piece{
		Shape: TPiece,
		Block: Purple,
		Points: []Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{0, 1},
		},
	}
	return p
}

func NewSPiece() *Piece {
	p := &Piece{
		Shape: SPiece,
		Block: Red,
		Points: []Coords{
			{0, 0},
			{0, 1},
			{1, 1},
			{1, 2},
		},
	}
	return p
}

func NewZPiece() *Piece {
	p := &Piece{
		Shape: ZPiece,
		Block: Yellow,
		Points: []Coords{
			{1, 0},
			{1, 1},
			{0, 1},
			{0, 2},
		},
	}
	return p
}

func NewJPiece() *Piece {
	p := &Piece{
		Shape: JPiece,
		Block: Green,
		Points: []Coords{
			{1, 0},
			{0, 1},
			{0, 0},
			{0, 2},
		},
	}
	return p
}
