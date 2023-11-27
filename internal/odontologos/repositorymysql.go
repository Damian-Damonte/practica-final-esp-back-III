package odontologos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (r *repositorymysql) GetAll(ctx context.Context) (*[]domain.Odontologo, error) {
	rows, err := r.db.Query(QueryGetAllOdontologos)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var odontologos []domain.Odontologo

	for rows.Next() {
		var odontologo domain.Odontologo
		err := rows.Scan(
			&odontologo.Id,
			&odontologo.Apellido,
			&odontologo.Nombre,
			&odontologo.Matricula,
		)

		if err != nil {
			return nil, err
		}

		odontologos = append(odontologos, odontologo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &odontologos, nil
}

func (r *repositorymysql) GetById(ctx context.Context, id int) (*domain.Odontologo, error) {
	row := r.db.QueryRow(QueryGetOdontologoById, id)

	var odontologo domain.Odontologo
	err := row.Scan(
		&odontologo.Id,
		&odontologo.Apellido,
		&odontologo.Nombre,
		&odontologo.Matricula,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		fmt.Println("error en repository GetById", err.Error())
		return nil, err
	}

	return &odontologo, nil
}

func (r *repositorymysql) Create(ctx context.Context, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	statement, err := r.db.Prepare(QuertyInsertOdontologo)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(odontologo.Apellido, odontologo.Nombre, odontologo.Matricula)

	if err != nil {
		return nil, ErrExecStatement
	}

	odontologoCreated := odontologo

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, ErrLastInsertedId
	}

	odontologoCreated.Id = int(lastId)

	return &odontologoCreated, nil
}

func (r *repositorymysql) Update(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	statement, err := r.db.Prepare(QueryUpdateOdontologo)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(odontologo.Apellido, odontologo.Nombre, odontologo.Matricula, id)
	if err != nil {
		return nil, ErrExecStatement
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}

	odontologoUpdated := odontologo
	odontologoUpdated.Id = id

	return &odontologoUpdated, nil
}

func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	result, err := r.db.Exec(QueryDeleteOdontologo, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

func (r *repositorymysql) Patch(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	panic("no implementado")
}
