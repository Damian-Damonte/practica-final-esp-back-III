package handlerpacientes

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/pacientes"
	"github.com/Damian-Damonte/practica-final-esp-back-III/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service pacientes.Service
}

func NewControladorPaciente(service pacientes.Repository) *Controlador {
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

		web.Success(ctx, http.StatusOK, gin.H{
			"data": pacientes,
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

		paciente, err := c.service.GetById(ctx, id)
		if err != nil {
			if errors.Is(err, pacientes.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d %s", "paciente con id", id, "no encontrado")
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": paciente,
		})
	}
}

func (c *Controlador) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pacienteReq domain.Paciente

		err := ctx.Bind(&pacienteReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		paciente, err := c.service.Create(ctx, pacienteReq)
		if err != nil {
			if errors.Is(err, pacientes.ErrPacienteAttributes) {
				web.Error(ctx, http.StatusBadRequest, "%s", err.Error())
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusCreated, gin.H{
			"data": paciente,
		})
	}
}

func (c *Controlador) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		var pacienteReq domain.Paciente

		err = ctx.Bind(&pacienteReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		pacienteUpdated, err := c.service.Update(ctx, id, pacienteReq)
		if err != nil {
			if errors.Is(err, pacientes.ErrPacienteAttributes) {
				web.Error(ctx, http.StatusBadRequest, "%s", err.Error())
				return
			}
			if errors.Is(err, pacientes.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d %s", "paciente con id", id, "no encontrado")
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": pacienteUpdated,
		})
	}
}

func (c *Controlador) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			if errors.Is(err, pacientes.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d %s", "paciente con id", id, "no encontrado")
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"message": fmt.Sprintf("paciente con id %d eliminado", id),
		})
	}
}