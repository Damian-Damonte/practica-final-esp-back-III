package domain

import (
	"time"
)

type Turno struct {
	Id          int        `json:"id"`
	Descripcion string     `json:"descripcion"`
	FechaHora   time.Time  `json:"fecha_hora"`
	Odontologo  Odontologo `json:"odontologo"`
	Paciente    Paciente   `json:"paciente"`
}

type TurnoMatriculaDni struct {
	Descripcion string    `json:"descripcion"`
	FechaHora   time.Time `json:"fecha_hora"`
	Matricula   string    `json:"matricula"`
	Dni         int       `json:"dni"`
}
