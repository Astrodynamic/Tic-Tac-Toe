package repository

import "sync"

type Storage struct {
	data sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Save(dto *DTO) {
	s.data.Store(dto.UUID, dto)
}

func (s *Storage) Load(uuid string) (*DTO, bool) {
	val, ok := s.data.Load(uuid)
	if !ok {
		return nil, false
	}
	return val.(*DTO), true
}
