package game

import (
	"image/color"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Grid        [][]*Sprite
	ActivePiece *Piece
	EmptyImage  *ebiten.Image
}

func NewBoard(row, col int) *Board {
	g := make([][]*Sprite, row)
	for i := range g {
		g[i] = make([]*Sprite, col)
	}
	e := ebiten.NewImage(40, 40)
	e.Fill(color.White)
	return &Board{
		Grid:       g,
		EmptyImage: e,
	}
}

func (b *Board) checkCollision(p *Piece) bool {
	g := b.Grid
	for _, coord := range p.Points {
		r := coord.R
		c := coord.C
		if r < 0 || r >= len(g) || c < 0 || c >= len(g[0]) || g[r][c] != nil {
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

func (b *Board) setPiece() {
	for _, v := range b.ActivePiece.Points {
		b.Grid[v.R][v.C] = b.ActivePiece.Sprite
	}
}

func (b *Board) applyGravity() bool {
	return b.movePiece(1, 0)
}

func (b *Board) instafall() {
	collide := false
	for !collide {
		collide = b.applyGravity()
	}
}

func (b *Board) deleteRow(row int) {
	for i := row; i > 0; i-- {
		for c := 0; c < len(b.Grid[i]); c++ {
			b.Grid[i][c] = b.Grid[i-1][c]
		}
	}
	for c := 0; c < len(b.Grid[0]); c++ {
		b.Grid[0][c] = nil
	}
}

func (b *Board) ClearLines() (deletedRowCount int) {
	deletedRowCount = 0
	rowsToCheck := make(map[int]bool)
	for _, v := range b.ActivePiece.Points {
		rowsToCheck[v.R] = true
	}
	rowsToDelete := []int{}
	b.setPiece()
	for row, _ := range rowsToCheck {
		fullRow := true
		for col := 0; col < len(b.Grid[row]); col++ {
			if b.Grid[row][col] == nil {
				fullRow = false
			}
		}
		if fullRow {
			deletedRowCount++
			rowsToDelete = append(rowsToDelete, row)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(rowsToDelete)))
	for i, row := range rowsToDelete {
		b.deleteRow(row + i)
	}
	return
}

func (b *Board) Draw(x, y int, screen *ebiten.Image) {
	for i, v := range b.Grid {
		for j, k := range v {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(x+(j*40)), float64(y+(i*40)))
			if k != nil {
				screen.DrawImage(k.Image, options)
			} else {
				screen.DrawImage(b.EmptyImage, options)
			}
		}
	}
	oldPiece := b.ActivePiece.copy()
	b.instafall()
	b.ActivePiece.DrawShadow(x, y, screen)
	b.ActivePiece = oldPiece
	b.ActivePiece.Draw(x, y, screen)
}
