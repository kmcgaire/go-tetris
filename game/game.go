package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	*Sprites
	NextPiece        *Piece
	Board            *Board
	gravityTickCount int
}

func (g *Game) Update() error {
	var keys []ebiten.Key
	g.gravityTickCount++
	for _, key := range inpututil.AppendPressedKeys(keys) {
		switch key {
		// For now only allow actions on first key press (will handle holding down keys later)
		case ebiten.KeyArrowUp:
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
				g.Board.rotatePiece()
			}
		case ebiten.KeyArrowDown:
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
				g.gravityTickCount = 0
				g.Board.movePiece(1, 0)
			}
		case ebiten.KeyArrowLeft:

			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				g.Board.movePiece(0, -1)
			}
		case ebiten.KeyArrowRight:
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				g.Board.movePiece(0, 1)
			}
		case ebiten.KeySpace:
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				g.gravityTickCount = 0
				g.Board.instafall()
				log.Printf("Deleted %d rows", g.Board.ClearLines())
				g.Board.ActivePiece = g.NextPiece
				g.NextPiece = GenerateRandomPiece(g.Sprites)
			}
		}
	}
	if g.gravityTickCount > 60 {
		if g.Board.applyGravity() {
			g.Board.ClearLines()
			g.Board.ActivePiece = g.NextPiece
			g.NextPiece = GenerateRandomPiece(g.Sprites)
		}
		g.gravityTickCount = 0
	}
	return nil
}

func NewGame(s *Sprites) *Game {
	b := NewBoard(20, 10)
	b.ActivePiece = GenerateRandomPiece(s)
	// Hack for now to ensure its on the screen
	b.ActivePiece.moveDown()
	b.ActivePiece.moveDown()
	b.ActivePiece.moveDown()
	b.ActivePiece.moveRight()
	b.ActivePiece.moveRight()
	return &Game{Sprites: s, NextPiece: GenerateRandomPiece(s), Board: b}
}

func (g *Game) DrawPiece(x, y int, screen *ebiten.Image, p *Piece) {
	block := g.blocks[p.Block]
	for _, v := range p.Points {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x+(v.C*40)), float64(y+(v.R*40)))
		screen.DrawImage(block.Image, options)
	}

}
func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawPiece(600, 200, screen, g.NextPiece)
	g.Board.Draw(50, 50, screen)
	ebitenutil.DebugPrintAt(screen, "Tetris V 0.0000010", 20, 20)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// log.Printf("Sprites in game update %v", g.Sprites)
	return outsideWidth, outsideHeight
}
