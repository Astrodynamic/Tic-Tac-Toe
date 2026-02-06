package domain

type CellType uint8

const (
	CellEmpty CellType = iota
	CellX
	CellO
)

type Board [BoardH][BoardW]CellType
