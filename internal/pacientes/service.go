package pacientes

import (
	"context"
	"errors"
	"time"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
)

var (
	ErrPacienteAttributes = errors.New("atributos de paciente incorrectos")
)

type service struct {
	repository Repository
}

func NewServicePaciente(repository Repository) Service {
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

func (s *service) Create(ctx context.Context, paciente domain.Paciente) (*domain.Paciente, error) {
	paciente.FechaAlta = time.Now()
	err := s.validatePaciente(paciente)
	if err != nil {
		return nil, err
	}

	pacienteCreated, err := s.repository.Create(ctx, paciente)
	if err != nil {
		return nil, err
	}

	return pacienteCreated, nil
}

func (s *service) Update(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error) {
	_, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	err = s.validatePaciente(paciente)
	if err != nil {
		return nil, err
	}

	pacienteUpdated, err := s.repository.Update(ctx, id, paciente)
	if err != nil {
		return nil, err
	}

	return pacienteUpdated, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Patch(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error) {
	pacienteSaved, err := s.repository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	pacientePatch, err := s.validatePacientePatch(*pacienteSaved, paciente)
	if err != nil {
		return nil, err
	}

	pacientePatched, err := s.repository.Patch(ctx, id, pacientePatch)
	if err != nil {
		return nil, err
	}

	return pacientePatched, nil
}

func (s *service) validatePaciente(paciente domain.Paciente) error {
	if paciente.Apellido == "" || paciente.Nombre == "" || paciente.Domicilio == "" ||
		paciente.Dni == 0 || paciente.FechaAlta.Equal(time.Time{}) {

		return ErrPacienteAttributes
	}

	return nil
}

func (s *service) validatePacientePatch(pacienteSaved, pacientePatch domain.Paciente) (domain.Paciente, error) {
	if pacientePatch.Apellido != "" {
		pacienteSaved.Apellido = pacientePatch.Apellido
	}
	if pacientePatch.Nombre != "" {
		pacienteSaved.Nombre = pacientePatch.Nombre
	}
	if pacientePatch.Domicilio != "" {
		pacienteSaved.Domicilio = pacientePatch.Domicilio
	}
	if pacientePatch.Dni != 0 {
		pacienteSaved.Dni = pacientePatch.Dni
	}
	if !pacientePatch.FechaAlta.Equal(time.Time{}) {
		pacienteSaved.FechaAlta = pacientePatch.FechaAlta
	}

	return pacienteSaved, nil
}
