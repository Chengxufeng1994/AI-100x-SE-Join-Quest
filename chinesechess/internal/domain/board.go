package domain

type Position struct {
	Row int
	Col int
}

type Piece struct {
	Name  string
	Color string
}

type Board struct {
	Pieces map[Position]Piece
}

func NewBoard() *Board {
	return &Board{
		Pieces: make(map[Position]Piece),
	}
}
