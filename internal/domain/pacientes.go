package domain

import "time"

type Paciente struct {
	Id int `json:"id"`
	Apellido  string    `json:"apellido"`
	Nombre    string    `json:"nombre"`
	Domicilio string    `json:"domicilio"`
	Dni       int       `json:"dni"`
	FechaAlta time.Time `json:"fecha_alta"`
}
