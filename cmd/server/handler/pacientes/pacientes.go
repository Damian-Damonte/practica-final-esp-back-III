package handlerpacientes

import (
	"net/http"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/pacientes"
	"github.com/Damian-Damonte/practica-final-esp-back-III/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service pacientes.Service
}

func NewControladorOdontologo(service pacientes.Repository) *Controlador {
	return &Controlador{
		service: service,
	}
}

func (c *Controlador) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pacientes, err := c.service.GetAll(ctx)

		if err != nil {
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H {
			"data": pacientes,
		})
	}
}
