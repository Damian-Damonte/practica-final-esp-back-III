package turnos

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement    = errors.New("error exec statement")
	ErrLastInsertedId   = errors.New("error last inserted id")
)

type repositorymysql struct {
	db *sql.DB
}

func NewMySqlRepository(db *sql.DB) Repository {
	return &repositorymysql{db: db}
}

func (r *repositorymysql) GetAll(ctx context.Context) (*[]domain.Turno, error) {
	rows, err := r.db.Query(QueryGetAllTurnos)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var turnos []domain.Turno

	for rows.Next() {
		var turno domain.Turno
		err := rows.Scan(
			&turno.Id,
			&turno.Descripcion,
			&turno.FechaHora,
			&turno.Odontologo.Id,
			&turno.Odontologo.Apellido,
			&turno.Odontologo.Nombre,
			&turno.Odontologo.Matricula,
			&turno.Paciente.Id,
			&turno.Paciente.Apellido,
			&turno.Paciente.Nombre,
			&turno.Paciente.Domicilio,
			&turno.Paciente.Dni,
			&turno.Paciente.FechaAlta,
		)

		if err != nil {
			return nil, err
		}

		turnos = append(turnos, turno)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &turnos, nil
}

func (r *repositorymysql) GetById(ctx context.Context, id int) (*domain.Turno, error) {
	row := r.db.QueryRow(QueryGetTurnoById, id)

	var turno domain.Turno
	err := row.Scan(
		&turno.Id,
			&turno.Descripcion,
			&turno.FechaHora,
			&turno.Odontologo.Id,
			&turno.Odontologo.Apellido,
			&turno.Odontologo.Nombre,
			&turno.Odontologo.Matricula,
			&turno.Paciente.Id,
			&turno.Paciente.Apellido,
			&turno.Paciente.Nombre,
			&turno.Paciente.Domicilio,
			&turno.Paciente.Dni,
			&turno.Paciente.FechaAlta,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &turno, nil
}
