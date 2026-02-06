package http

import "tictactoe/internal/domain"

func toResponse(g *domain.Game, winner domain.CellType, over bool) *Response {
	resp := &Response{UUID: g.UUID}
	for i := range g.Board {
		for j := range g.Board[i] {
			resp.Board[i][j] = uint8(g.Board[i][j])
		}
	}

	if over {
		w := uint8(winner)
		resp.Winner = &w
	}

	return resp
}

func fromRequest(req *Request) domain.Board {
	var board domain.Board
	for i := range req.Board {
		for j := range req.Board[i] {
			board[i][j] = domain.CellType(req.Board[i][j])
		}
	}
	return board
}
