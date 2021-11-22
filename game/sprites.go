package game

import (
	"image"
	"image/png"
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

type Block int

const (
	Empty Block = iota
	LightBlue
	Blue
	Pink
	Purple
	Red
	Yellow
	Green
	Grey
)

type Sprite struct {
	Block
	Image *ebiten.Image
}

type Sprites struct {
	blocks map[Block]*Sprite
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func (s *Sprites) Block(b Block) *ebiten.Image {
	sprite := s.blocks[b]
	if sprite.Image == nil {
		log.Fatalf("invalid image name: %d", b)
	}
	return sprite.Image
}

func LoadSprites() (*Sprites, error) {
	s := &Sprites{
		blocks: make(map[Block]*Sprite),
	}

	file, err := os.Open(blocksPath)
	if err != nil {
		log.Printf("Failed to open image %v", err)
		return nil, err
	}

	defer file.Close()

	// Load Image
	img, err := png.Decode(file)
	if err != nil {
		log.Printf("Failed to decode image %v", err)
		return nil, err
	}

	for i := 0; i < 16; i++ {
		r := i / cols
		c := i % cols
		subImage := img.(SubImager).SubImage(image.Rect(c*tileSize, r*tileSize, (c+1)*tileSize, (r+1)*tileSize))
		s.blocks[Block(i+1)] = &Sprite{Block(i + 1), ebiten.NewImageFromImage(subImage)}
	}
	log.Printf("Image Bounds: %v", img.Bounds())
	return s, nil
}
