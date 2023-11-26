package domain

import "time"

type Turno struct {
	Id           int       `json:"id"`
	Descripcion  string    `json:"descripcion"`
	FechaHora    time.Time `json:"fecha_hora"`
	OdontologoId int       `json:"odontologo_id"`
	PacienteId   int       `json:"paciente_id"`
}
