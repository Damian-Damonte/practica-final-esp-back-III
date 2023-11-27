package odontologos

import (
	"context"
	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
)

type service struct {
	repository Repository
}

func NewServiceOdontologo(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll(ctx context.Context) (*[]domain.Odontologo, error) {
	listOdontologos, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return listOdontologos, nil
}

func (s *service) GetById(ctx context.Context, id int) (*domain.Odontologo, error) {
	odontologo, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return odontologo, nil
}

func (s *service) Create(ctx context.Context, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	panic("no implementado")
}

func (s *service) Update(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	panic("no implementado")
}

func (s *service) Delete(ctx context.Context, id int) error {
	panic("no implementado")
}

func (s *service) Patch(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	panic("no implementado")
}
