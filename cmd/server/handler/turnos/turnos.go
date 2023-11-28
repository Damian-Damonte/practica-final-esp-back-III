package handlerturnos

import (
	"net/http"
	"strconv"
	"errors"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/turnos"
	"github.com/Damian-Damonte/practica-final-esp-back-III/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service turnos.Service
}

func NewControladorOdontologo(service turnos.Repository) *Controlador {
	return &Controlador{
		service: service,
	}
}

func (c *Controlador) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		turnos, err := c.service.GetAll(ctx)

		if err != nil {
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turnos,
		})
	}
}

func (c *Controlador) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		turno, err := c.service.GetById(ctx, id)
		if err != nil {
			if errors.Is(err, turnos.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d %s", "turno con id", id, "no encontrado")
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turno,
		})
	}
}

func (c *Controlador) HandlerGetPacienteDni() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dni, err := strconv.Atoi(ctx.Query("dni"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "dni invalido")
			return
		}

		turnosPaciente, err := c.service.GetByPacienteDni(ctx, dni)
		if err != nil {
			if errors.Is(err, turnos.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d", "no se encontraron turnos para el paciente con dni ", dni)
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": turnosPaciente,
		})
	}
}