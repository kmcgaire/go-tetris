package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	*Sprites
	NextPiece *Piece
}

func (g *Game) Update() error {
	var keys []ebiten.Key
	for _, key := range inpututil.AppendPressedKeys(keys) {
		switch key {
		// For now only allow actions on first key press (will handle holding down keys later)
		case ebiten.KeyArrowUp:
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
				g.NextPiece = g.NextPiece.rotate()
			}
		case ebiten.KeyArrowDown:
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
				g.NextPiece = g.NextPiece.moveDown()
			}
		case ebiten.KeyArrowLeft:

			if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
				g.NextPiece = g.NextPiece.moveLeft()
			}
		case ebiten.KeyArrowRight:

			if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
				g.NextPiece = g.NextPiece.moveRight()
			}
		}
	}
	return nil
}

func NewGame(s *Sprites) *Game {

	return &Game{Sprites: s, NextPiece: GenerateRandomPiece()}
}

func (g *Game) DrawPiece(x, y int, screen *ebiten.Image, p *Piece) {
	block := g.blocks[int(p.Block)]
	for _, v := range p.Points {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x+(v.C*40)), float64(y+(v.R*40)))
		screen.DrawImage(block, options)
	}

}
func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawPiece(500, 200, screen, g.NextPiece)
	ebitenutil.DebugPrintAt(screen, "Tetris V 0.0000010", 20, 20)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// log.Printf("Sprites in game update %v", g.Sprites)
	return outsideWidth, outsideHeight
}
