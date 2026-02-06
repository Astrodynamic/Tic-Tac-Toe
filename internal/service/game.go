package service

import "tictactoe/internal/domain"

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) domain.Service {
	return &Service{repo: repo}
}

func (s *Service) Move(b domain.Board) domain.Board {
	_, p := s.best(b, domain.CellO)
	if p.r >= 0 {
		b[p.r][p.c] = domain.CellO
	}
	return b
}

func (s *Service) Validate(g *domain.Game, nb domain.Board) error {
	ob := domain.Board{}
	if g != nil {
		ob = g.Board
	}

	diff := 0
	for i := range nb {
		for j := range nb[i] {
			if ob[i][j] != nb[i][j] {
				if ob[i][j] != domain.CellEmpty {
					if ob[i][j] == domain.CellO || nb[i][j] != domain.CellX {
						return domain.ErrBoardAltered
					}
				}
				diff++
			}
		}
	}

	if diff == 0 {
		return domain.ErrNoMoveMade
	}
	if diff > 1 {
		return domain.ErrTooManyMoves
	}
	return nil
}

func (s *Service) Done(b domain.Board) (bool, domain.CellType) {
	if w := s.winner(b); w != domain.CellEmpty {
		return true, w
	}
	if s.full(b) {
		return true, domain.CellEmpty
	}
	return false, domain.CellEmpty
}

func (s *Service) Save(g *domain.Game) error {
	return s.repo.Save(g)
}

func (s *Service) Load(uuid string) (*domain.Game, error) {
	if g, err := s.repo.Load(uuid); err == nil {
		return g, nil
	}
	g := &domain.Game{UUID: uuid}
	return g, s.repo.Save(g)
}
