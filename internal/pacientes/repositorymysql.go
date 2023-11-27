package pacientes

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

func (r *repositorymysql) GetAll(ctx context.Context) (*[]domain.Paciente, error) {
	rows, err := r.db.Query(QueryGetAllPacientes)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pacientes []domain.Paciente

	for rows.Next() {
		var paciente domain.Paciente
		err := rows.Scan(
			&paciente.Id,
			&paciente.Apellido,
			&paciente.Nombre,
			&paciente.Domicilio,
			&paciente.Dni,
			&paciente.FechaAlta,
		)

		if err != nil {
			return nil, err
		}

		pacientes = append(pacientes, paciente)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pacientes, nil
}

func (r *repositorymysql) GetById(ctx context.Context, id int) (*domain.Paciente, error) {
	row := r.db.QueryRow(QueryGetPacientesById, id)

	var paciente domain.Paciente
	err := row.Scan(
		&paciente.Id,
		&paciente.Apellido,
		&paciente.Nombre,
		&paciente.Domicilio,
		&paciente.Dni,
		&paciente.FechaAlta,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &paciente, nil
}

func (r *repositorymysql) Create(ctx context.Context, paciente domain.Paciente) (*domain.Paciente, error) {
	statement, err := r.db.Prepare(QueryInsertPaciente)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(paciente.Apellido, paciente.Nombre, paciente.Domicilio, paciente.Dni, paciente.FechaAlta)

	if err != nil {
		return nil, ErrExecStatement
	}

	pacienteCreadte := paciente

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, ErrLastInsertedId
	}

	pacienteCreadte.Id = int(lastId)

	return &pacienteCreadte, nil
}

func (r *repositorymysql) Update(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error) {
	statement, err := r.db.Prepare(QueryUpdatePaciente)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(paciente.Apellido, paciente.Nombre, paciente.Domicilio, paciente.Dni, paciente.FechaAlta, id)
	if err != nil {
		return nil, ErrExecStatement
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}

	pacienteUpdated := paciente
	pacienteUpdated.Id = id

	return &pacienteUpdated, nil
}

func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	panic("no implementado")
}

func (r *repositorymysql) Patch(ctx context.Context, id int, paciente domain.Paciente) (*domain.Paciente, error) {
	panic("no implementado")
}
