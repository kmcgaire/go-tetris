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
	Points [4]Coords
}

func NewLPiece() *Piece {
	p := &Piece{
		Shape: LPiece,
		Block: LightBlue,
		Points: [4]Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{0, 0},
		},
	}
	return p
}

func NewIPiece() *Piece {
	p := &Piece{
		Shape: IPiece,
		Block: Blue,
		Points: [4]Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{1, 3},
		},
	}
	return p
}

func NewOPiece() *Piece {
	p := &Piece{
		Shape: OPiece,
		Block: Pink,
		Points: [4]Coords{
			{1, 0},
			{1, 1},
			{0, 0},
			{0, 1},
		},
	}
	return p
}

func NewTPiece() *Piece {
	p := &Piece{
		Shape: TPiece,
		Block: Purple,
		Points: [4]Coords{
			{0, 0},
			{0, 1},
			{1, 1},
			{1, 2},
		},
	}
	return p
}

func NewSPiece() *Piece {
	p := &Piece{
		Shape: SPiece,
		Block: Red,
		Points: [4]Coords{
			{1, 0},
			{1, 1},
			{0, 1},
			{0, 2},
		},
	}
	return p
}

func NewZPiece() *Piece {
	p := &Piece{
		Shape: ZPiece,
		Block: Yellow,
		Points: [4]Coords{
			{1, 0},
			{0, 1},
			{0, 0},
			{0, 2},
		},
	}
	return p
}

func NewJPiece() *Piece {
	p := &Piece{
		Shape: JPiece,
		Block: Green,
		Points: [4]Coords{
			{1, 0},
			{1, 1},
			{1, 2},
			{0, 0},
		},
	}
	return p
}
