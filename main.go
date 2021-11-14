package main

import (
	"log"
	"tetris/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := &game.Game{}

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(680, 480)
	ebiten.SetWindowTitle("Tetris Lets goooo")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
