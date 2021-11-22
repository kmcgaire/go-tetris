package game

// Piece is a constant for a shape of piece. There are 7 classic pieces like L, and O
type Shape int

// Various values that the pieces can be
const (
	IPiece Shape = iota
	JPiece
	LPiece
	OPiece
	SPiece
	TPiece
	ZPiece
)

type Coords struct {
	R, C int
}

type Piece struct {
	Shape
	Block
	points [4]Coords
}

func NewLPiece() {

}
