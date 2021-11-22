package game

import (
	"math"

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

func DrawPiece(x, y int, screen, block *ebiten.Image) {
	x = 50 + x*40
	y = 50 + y*40
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(block, options)
}
func (g *Game) Draw(screen *ebiten.Image) {
	// TODO(kevin) Actually draw things properly but this looks pretty
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			block := g.blocks[int(math.Mod(float64(i+j), 15)+1)]
			DrawPiece(i, j, screen, block)
		}
	}

	// Test drawing purple
	for i := 0; i < 10; i++ {
		DrawPiece(i, 19, screen, g.blocks[int(Purple)])
	}

	ebitenutil.DebugPrintAt(screen, "Tetris V 0.0000003", 20, 20)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// log.Printf("Sprites in game update %v", g.Sprites)
	return outsideWidth, outsideHeight
}
