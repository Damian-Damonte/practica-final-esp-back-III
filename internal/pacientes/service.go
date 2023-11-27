package pacientes

import (
	"context"
	"errors"
	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
)

var (
	ErrPacienteUpdate = errors.New("atributos de paciente incorrectos")
)

type service struct {
	repository Repository
}

func NewServiceOdontologo(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetAll(ctx context.Context) (*[]domain.Paciente, error) {
	listPacientes, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return listPacientes, nil
}

func (s *service) GetById(ctx context.Context, id int) (*domain.Paciente, error) {
	paciente, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return paciente, nil
}

func (s *service) Create(ctx context.Context, odontologo domain.Paciente) (*domain.Paciente, error) {
	panic("no implementado")
}

func (s *service) Update(ctx context.Context, id int, odontologo domain.Paciente) (*domain.Paciente, error) {
	panic("no implementado")
}

func (s *service) Delete(ctx context.Context, id int) error {
	panic("no implementado")
}

func (s *service) Patch(ctx context.Context, id int, odontologo domain.Paciente) (*domain.Paciente, error) {
	panic("no implementado")
}
