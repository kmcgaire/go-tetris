package game

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	blocksPath = "game/sprites/blocks.png"
	rows       = 2
	cols       = 8
	tileSize   = 40
)

type Sprites struct {
	blocks map[int]*ebiten.Image
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func (s *Sprites) Block(i int) *ebiten.Image {
	image := s.blocks[i]
	if image == nil {
		log.Fatalf("invalid image name: %d", i)
	}
	return image
}

func LoadSprites() (*Sprites, error) {
	s := &Sprites{
		blocks: make(map[int]*ebiten.Image),
	}

	file, err := os.Open(blocksPath)
	if err != nil {
		log.Printf("Failed to open image %v", err)
		return nil, err
	}

	defer file.Close()

	// Load Image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Printf("Failed to decode image %v", err)
		return nil, err
	}

	for i := 0; i < 16; i++ {
		r := i / cols
		c := i % cols
		subImage := img.(SubImager).SubImage(image.Rect(c*tileSize, r*tileSize, (c+1)*tileSize, (r+1)*tileSize))
		s.blocks[i] = ebiten.NewImageFromImage(subImage)
	}
	log.Printf("Image Bounds: %v", img.Bounds())
	return s, nil
}
