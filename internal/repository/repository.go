package repository

import "tictactoe/internal/domain"

type Repository struct {
	storage *Storage
}

func NewRepository(storage *Storage) domain.Repository {
	return &Repository{storage: storage}
}

func (r *Repository) Save(g *domain.Game) error {
	r.storage.Save(toDTO(g))
	return nil
}

func (r *Repository) Load(uuid string) (*domain.Game, error) {
	dto, ok := r.storage.Load(uuid)
	if !ok {
		return nil, domain.ErrGameNotFound
	}
	return fromDTO(dto), nil
}
