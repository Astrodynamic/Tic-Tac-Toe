package domain

type Service interface {
	Move(board Board) Board
	Validate(game *Game, board Board) error
	Done(board Board) (bool, CellType)
	Save(game *Game) error
	Load(uuid string) (*Game, error)
}

type Repository interface {
	Save(game *Game) error
	Load(uuid string) (*Game, error)
}
