package game

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	*Sprites
	NextPiece        *Piece
	Board            *Board
	gravityTickCount float64
	score            int
	level            int
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
				g.clearLines()
			}
		}
	}
	if g.gravityTickCount > math.Max(float64(60-(g.level*7)), 15) {
		if g.Board.applyGravity() {
			g.clearLines()
		}
		g.gravityTickCount = 0
	}
	return nil
}

func (g *Game) clearLines() {
	linesCleared := g.Board.ClearLines()
	// Bonus points for clearing more lines
	if linesCleared > 0 {
		g.score += (linesCleared*200 + (linesCleared-1)*300) * g.level
		if g.score >= 10000 {
			g.level = 8
		} else if g.score >= 5000 {
			g.level = 7
		} else if g.score >= 3000 {
			g.level = 6
		} else if g.score >= 2500 {
			g.level = 5
		} else if g.score >= 1000 {
			g.level = 4
		} else if g.score >= 800 {
			g.level = 3
		} else if g.score >= 300 {
			g.level = 2
		}
	}
	g.Board.ActivePiece = g.NextPiece
	g.NextPiece = GenerateRandomPiece(g.Sprites)
}

func NewGame(s *Sprites) *Game {
	b := NewBoard(20, 10)
	b.ActivePiece = GenerateRandomPiece(s)
	return &Game{Sprites: s, NextPiece: GenerateRandomPiece(s), Board: b, level: 1}
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.NextPiece.Draw(600, 200, screen)
	g.Board.Draw(50, 50, screen)
	ebitenutil.DebugPrintAt(screen, "Tetris V 0.0000010", 20, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Score: %d", g.score), 700, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("level: %d", g.level), 700, 40)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// log.Printf("Sprites in game update %v", g.Sprites)
	return outsideWidth, outsideHeight
}
