package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	*Sprites
}

func (g *Game) Update() error {
	return nil
}

func NewGame(s *Sprites) *Game {
	return &Game{Sprites: s}
}

func (g *Game) DrawPiece(x, y int, screen *ebiten.Image, p *Piece) {
	block := g.blocks[int(p.Block)]
	for _, v := range p.Points {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(x+(v.R*40)), float64(y+(v.C*40)))
		screen.DrawImage(block, options)
	}

}
func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawPiece(50, 50, screen, NewLPiece())
	g.DrawPiece(50*2, 50, screen, NewIPiece())
	g.DrawPiece(50*4, 50, screen, NewOPiece())
	g.DrawPiece(50*8, 50, screen, NewTPiece())
	g.DrawPiece(50, 50*4, screen, NewSPiece())
	g.DrawPiece(50, 50*8, screen, NewZPiece())
	g.DrawPiece(50*4, 50*8, screen, NewJPiece())
	ebitenutil.DebugPrintAt(screen, "Tetris V 0.0000003", 20, 20)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// log.Printf("Sprites in game update %v", g.Sprites)
	return outsideWidth, outsideHeight
}
