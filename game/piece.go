package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
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
	*Sprite
	Points []Coords
}

func (p *Piece) copy() *Piece {
	np := Piece{Shape: p.Shape, Sprite: p.Sprite, Points: make([]Coords, 4)}
	copy(np.Points, p.Points)
	return &np
}

func (p *Piece) Draw(x, y int, screen *ebiten.Image) {
	for _, v := range p.Points {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x+(v.C*40)), float64(y+(v.R*40)))
		screen.DrawImage(p.Image, options)
	}
}

func (p *Piece) rotate() *Piece {
	// OPiece shouldn't be rotated
	if p.Shape == OPiece {
		return p
	}
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

func GenerateRandomPiece(s *Sprites) *Piece {
	i := rand.Intn(7)
	switch i {
	case 0:
		return NewLPiece(s)
	case 1:
		return NewIPiece(s)
	case 2:
		return NewOPiece(s)
	case 3:
		return NewTPiece(s)
	case 4:
		return NewSPiece(s)
	case 5:
		return NewZPiece(s)
	default:
		return NewJPiece(s)
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

func NewLPiece(s *Sprites) *Piece {
	p := &Piece{
		Shape:  LPiece,
		Sprite: &Sprite{LightBlue, s.blocks[LightBlue].Image},
		Points: []Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{0, 0},
		},
	}
	return p
}

func NewIPiece(s *Sprites) *Piece {
	p := &Piece{
		Shape:  IPiece,
		Sprite: &Sprite{Blue, s.blocks[Blue].Image},
		Points: []Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{1, 3},
		},
	}
	return p
}

func NewOPiece(s *Sprites) *Piece {
	p := &Piece{
		Shape:  OPiece,
		Sprite: &Sprite{Pink, s.blocks[Pink].Image},
		Points: []Coords{
			{1, 0},
			{1, 1},
			{0, 0},
			{0, 1},
		},
	}
	return p
}

func NewTPiece(s *Sprites) *Piece {
	p := &Piece{
		Shape:  TPiece,
		Sprite: &Sprite{Purple, s.blocks[Purple].Image},
		Points: []Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{0, 1},
		},
	}
	return p
}

func NewSPiece(s *Sprites) *Piece {
	p := &Piece{
		Shape:  SPiece,
		Sprite: &Sprite{Red, s.blocks[Red].Image},
		Points: []Coords{
			{0, 0},
			{0, 1},
			{1, 1},
			{1, 2},
		},
	}
	return p
}

func NewZPiece(s *Sprites) *Piece {
	p := &Piece{
		Shape:  ZPiece,
		Sprite: &Sprite{Yellow, s.blocks[Yellow].Image},
		Points: []Coords{
			{1, 0},
			{1, 1},
			{0, 1},
			{0, 2},
		},
	}
	return p
}

func NewJPiece(s *Sprites) *Piece {
	p := &Piece{
		Shape:  JPiece,
		Sprite: &Sprite{Green, s.blocks[Green].Image},
		Points: []Coords{
			{1, 0},
			{0, 1},
			{0, 0},
			{0, 2},
		},
	}
	return p
}
