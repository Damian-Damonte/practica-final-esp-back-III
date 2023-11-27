package pacientes

import (
	"context"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) (*[]domain.Paciente, error)
	GetById(ctx context.Context, id int) (*domain.Paciente, error)
	Create(ctx context.Context, odontologo domain.Paciente) (*domain.Paciente, error)
	Update(ctx context.Context, id int, odontologo domain.Paciente) (*domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, odontologo domain.Paciente) (*domain.Paciente, error)
}

type Service interface {
	GetAll(ctx context.Context) (*[]domain.Paciente, error)
	GetById(ctx context.Context, id int) (*domain.Paciente, error)
	Create(ctx context.Context, odontologo domain.Paciente) (*domain.Paciente, error)
	Update(ctx context.Context, id int, odontologo domain.Paciente) (*domain.Paciente, error)
	Delete(ctx context.Context, id int) error
	Patch(ctx context.Context, id int, odontologo domain.Paciente) (*domain.Paciente, error)
}