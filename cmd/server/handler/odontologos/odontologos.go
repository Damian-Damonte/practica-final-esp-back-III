package handlerodontologos

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/domain"
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

		web.Success(ctx, http.StatusOK, gin.H{
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

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologo,
		})
	}
}

func (c *Controlador) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var odontologoReq domain.Odontologo

		err := ctx.Bind(&odontologoReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		odontologo, err := c.service.Create(ctx, odontologoReq)
		if err != nil {
			if errors.Is(err, odontologos.ErrOdontologoAttributes) {
				web.Error(ctx, http.StatusBadRequest, "%s", err.Error())
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusCreated, gin.H{
			"data": odontologo,
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

		var odontologoReq domain.Odontologo

		err = ctx.Bind(&odontologoReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		odontologoUpdated, err := c.service.Update(ctx, id, odontologoReq)
		if err != nil {
			if errors.Is(err, odontologos.ErrOdontologoAttributes) {
				web.Error(ctx, http.StatusBadRequest, "%s", err.Error())
				return
			}
			if errors.Is(err, odontologos.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d %s", "odontologo con id", id, "no encontrado")
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologoUpdated,
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
			if errors.Is(err, odontologos.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d %s", "odontologo con id", id, "no encontrado")
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"message": fmt.Sprintf("odontologo con id %d eliminado", id),
		})
	}
}

func (c *Controlador) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		var odontologoReq domain.Odontologo

		err = ctx.Bind(&odontologoReq)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		odontologoPatched, err := c.service.Patch(ctx, id, odontologoReq)
		if err != nil {
			if errors.Is(err, odontologos.ErrNotFound) {
				web.Error(ctx, http.StatusNotFound, "%s %d %s", "odontologo con id", id, "no encontrado")
				return
			}
			web.InternalServerError(ctx)
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologoPatched,
		})
	}
}
