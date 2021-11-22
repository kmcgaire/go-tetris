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
	pieces := []*Piece{
		NewLPiece(),
		NewIPiece(),
		NewOPiece(),
		NewTPiece(),
		NewSPiece(),
		NewZPiece(),
		NewJPiece(),
	}
	for i, p := range pieces {
		g.DrawPiece(40*2*i, 50, screen, p)
	}
	ebitenutil.DebugPrintAt(screen, "Tetris V 0.0000003", 20, 20)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// log.Printf("Sprites in game update %v", g.Sprites)
	return outsideWidth, outsideHeight
}
