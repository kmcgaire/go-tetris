package main

import (
	"log"
	"tetris/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	s, err := game.LoadSprites()
	if err != nil {
		log.Fatalf("Could not generate sprites %v", err)
		return
	}
	game := game.NewGame(s)

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Tetris Lets goooo")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
