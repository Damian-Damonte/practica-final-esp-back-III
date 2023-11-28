package turnos

import (
	"context"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) (*[]domain.Turno, error)
	GetById(ctx context.Context, id int) (*domain.Turno, error)
	// GetByDni(ctx context.Context, dni int) (*[]domain.Turno, error)
	// CreateByMatriculaAndDni(ctx context.Context, turnoMatriculaDni domain.TurnoMatriculaDni) (*domain.Turno, error)
	// Update(ctx context.Context, id int, paciente domain.Turno) (*domain.Turno, error)
	// Delete(ctx context.Context, id int) error
	// Patch(ctx context.Context, id int, paciente domain.Turno) (*domain.Turno, error)
}

type Service interface {
	GetAll(ctx context.Context) (*[]domain.Turno, error)
	GetById(ctx context.Context, id int) (*domain.Turno, error)
	// GetByDni(ctx context.Context, dni int) (*[]domain.Turno, error)
	// Create(ctx context.Context, paciente domain.Turno) (*domain.Turno, error)
	// Update(ctx context.Context, id int, paciente domain.Turno) (*domain.Turno, error)
	// Delete(ctx context.Context, id int) error
	// Patch(ctx context.Context, id int, paciente domain.Turno) (*domain.Turno, error)
}