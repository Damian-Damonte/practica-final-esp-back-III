package odontologos

import (
	"context"
	"database/sql"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
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
	panic("no implementado")
}

func (r *repositorymysql) Create(ctx context.Context, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	panic("no implementado")
}

func (r *repositorymysql) Update(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	panic("no implementado")
}

func (r *repositorymysql) Delete(ctx context.Context, id int) error {
	panic("no implementado")
}

func (r *repositorymysql) Patch(ctx context.Context, id int, odontologo domain.Odontologo) (*domain.Odontologo, error) {
	panic("no implementado")
}
