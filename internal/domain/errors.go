package domain

import "errors"

var (
	ErrGameNotFound    = errors.New("game not found")
	ErrInvalidBoard    = errors.New("invalid board")
	ErrBoardAltered    = errors.New("previous moves have been altered")
	ErrInvalidMove     = errors.New("invalid move")
	ErrGameAlreadyOver = errors.New("game is already over")
	ErrTooManyMoves    = errors.New("too many moves made")
	ErrNoMoveMade      = errors.New("no move made")
)
