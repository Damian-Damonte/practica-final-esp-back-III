package handlerodontologos

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/odontologos"
	"github.com/Damian-Damonte/practica-final-esp-back-III/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controlador struct {
	service odontologos.Service
}

func NewControladorOdontologo(service odontologos.Repository) *Controlador {
	return &Controlador{
		service: service,
	}
}

func (c *Controlador) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		odontologos, err := c.service.GetAll(ctx)

		if err != nil {
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H {
			"data": odontologos,
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

		odontologo, err := c.service.GetById(ctx, id)
		if err != nil {
			if errors.Is(err, odontologos.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d %s", "odontologo con id", id, "no encontrado")
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H {
			"data": odontologo,
		})
	}
}