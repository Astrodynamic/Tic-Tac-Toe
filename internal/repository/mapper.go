package repository

import "tictactoe/internal/domain"

func toDTO(g *domain.Game) *DTO {
	dto := &DTO{UUID: g.UUID}
	for i := range g.Board {
		for j := range g.Board[i] {
			dto.Board[i][j] = uint8(g.Board[i][j])
		}
	}
	return dto
}

func fromDTO(dto *DTO) *domain.Game {
	g := &domain.Game{UUID: dto.UUID}
	for i := range dto.Board {
		for j := range dto.Board[i] {
			g.Board[i][j] = domain.CellType(dto.Board[i][j])
		}
	}
	return g
}
