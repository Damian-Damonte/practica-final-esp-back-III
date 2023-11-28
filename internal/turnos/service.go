package turnos

import (
	"context"
	"errors"
	"fmt"

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
		fmt.Println(err)
		return nil, err
	}

	return listTurnos, nil
}
