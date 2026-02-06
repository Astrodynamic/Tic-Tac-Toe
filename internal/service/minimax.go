package service

import "tictactoe/internal/domain"

const (
	win = domain.BoardW*domain.BoardH + 1
	inf = win + 1
)

type pos struct{ r, c int }

func (s *Service) best(b domain.Board, player domain.CellType) (int, pos) {
	score := -inf
	p := pos{-1, -1}

	for i := range b {
		for j := range b[i] {
			if b[i][j] == domain.CellEmpty {
				b[i][j] = player
				s := s.minimax(b, 0, false)
				b[i][j] = domain.CellEmpty
				if s > score {
					score = s
					p = pos{i, j}
				}
			}
		}
	}
	return score, p
}

func (s *Service) minimax(b domain.Board, d int, max bool) int {
	if w := s.winner(b); w != domain.CellEmpty {
		return s.score(w, d)
	}
	if s.full(b) {
		return 0
	}
	return s.eval(b, d, max)
}

func (s *Service) score(w domain.CellType, d int) int {
	if w == domain.CellO {
		return win - d
	}
	return d - win
}

func (s *Service) eval(b domain.Board, d int, max bool) int {
	best := inf
	p := domain.CellX
	cmp := func(a, b int) bool { return a < b }

	if max {
		best = -inf
		p = domain.CellO
		cmp = func(a, b int) bool { return a > b }
	}

	for i := range b {
		for j := range b[i] {
			if b[i][j] == domain.CellEmpty {
				b[i][j] = p
				s := s.minimax(b, d+1, !max)
				b[i][j] = domain.CellEmpty
				if cmp(s, best) {
					best = s
				}
			}
		}
	}
	return best
}

func (s *Service) winner(b domain.Board) domain.CellType {
	lines := [][3]pos{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, ln := range lines {
		c := b[ln[0].r][ln[0].c]
		if c != domain.CellEmpty && c == b[ln[1].r][ln[1].c] && c == b[ln[2].r][ln[2].c] {
			return c
		}
	}
	return domain.CellEmpty
}

func (s *Service) full(b domain.Board) bool {
	for i := range b {
		for j := range b[i] {
			if b[i][j] == domain.CellEmpty {
				return false
			}
		}
	}
	return true
}
