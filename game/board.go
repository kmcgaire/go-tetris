package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Grid        [][]Sprite
	ActivePiece *Piece
}

func NewBoard(row, col int) *Board {
	g := make([][]Sprite, row)
	for i := range g {
		g[i] = make([]Sprite, col)
	}
	return &Board{
		Grid: g,
	}
}

func (b *Board) checkCollision(p *Piece) bool {
	g := b.Grid
	for _, coord := range p.Points {
		r := coord.R
		c := coord.C
		log.Printf("Checking collision r: %v c: %v", r, c)
		if r < 0 || r >= len(g) || c < 0 || c >= len(g[0]) || g[r][c].Block != Empty {
			return true
		}
	}
	return false
}

func (b *Board) rotatePiece() {
	newPiece := b.ActivePiece.rotate()
	if b.checkCollision(newPiece) {
		if !b.checkCollision(newPiece.moveRight()) {
			b.ActivePiece = newPiece.moveRight()
		} else if !b.checkCollision(newPiece.moveLeft()) {
			b.ActivePiece = newPiece.moveLeft()
		} else if !b.checkCollision(newPiece.moveDown()) {
			b.ActivePiece = newPiece.moveDown()
		}
	} else {
		b.ActivePiece = newPiece
	}
}

func (b *Board) movePiece(r, c int) bool {
	newPiece := b.ActivePiece.move(r, c)
	if !b.checkCollision(newPiece) {
		b.ActivePiece = newPiece
		return false
	}
	return true
}

func (b *Board) applyGravity() bool {
	didCollide := b.movePiece(1, 0)
	if didCollide {
		// Set the current block in the grid
		for _, v := range b.ActivePiece.Points {
			b.Grid[v.R][v.C] = *b.ActivePiece.Sprite
		}
		b.ActivePiece = nil
		return true
	}
	return false
}

func (b *Board) instafall() {
	collide := false
	for !collide {
		collide = b.applyGravity()
	}
}

func (b *Board) Draw(x, y int, screen *ebiten.Image) {
	for i, v := range b.Grid {
		for j, k := range v {
			if k.Block != Empty {
				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(x+(j*40)), float64(y+(i*40)))
				screen.DrawImage(k.Image, options)
			}
		}
	}
	for _, v := range b.ActivePiece.Points {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x+(v.C*40)), float64(y+(v.R*40)))
		screen.DrawImage(b.ActivePiece.Image, options)
	}
}
