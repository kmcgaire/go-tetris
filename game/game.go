package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	*Sprites
	Pieces []*Piece
}

func (g *Game) Update() error {
	var keys []ebiten.Key
	for _, key := range inpututil.AppendPressedKeys(keys) {
		switch key {
		case ebiten.KeyArrowUp:
			// Only allow rotation on first key press
			if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
				for i, p := range g.Pieces {
					g.Pieces[i] = p.rotate()
				}
			}
			// When movement comes allow some sort key hold to move probably?
		}
	}
	return nil
}

func NewGame(s *Sprites) *Game {
	return &Game{Sprites: s, Pieces: []*Piece{
		NewLPiece(),
		NewIPiece(),
		NewOPiece(),
		NewTPiece(),
		NewSPiece(),
		NewZPiece(),
		NewJPiece(),
	}}
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
	for i, p := range g.Pieces {
		g.DrawPiece(40*3*i, 40*3*i, screen, p)
	}
	ebitenutil.DebugPrintAt(screen, "Tetris V 0.0000003", 20, 20)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// log.Printf("Sprites in game update %v", g.Sprites)
	return outsideWidth, outsideHeight
}
