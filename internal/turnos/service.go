package turnos

import (
	"context"
	"errors"
	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
)

var (
	ErrTurnoAttributes = errors.New("atributos de turno incorrectos")
)

type service struct {
	repository Repository
}

func NewServiceTurno(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll(ctx context.Context) (*[]domain.Turno, error) {
	listTurnos, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return listTurnos, nil
}

func (s *service) GetById(ctx context.Context, id int) (*domain.Turno, error) {
	paciente, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return paciente, nil
}

func (s *service) GetByPacienteDni(ctx context.Context, dni int) (*[]domain.Turno, error) {
	listTurnos, err := s.repository.GetByPacienteDni(ctx, dni)
	if err != nil {
		return nil, err
	}

	return listTurnos, nil
}
