package handlerodontologos

import (
	"net/http"

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
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, gin.H{
			"data": odontologos,
		})
	}
}