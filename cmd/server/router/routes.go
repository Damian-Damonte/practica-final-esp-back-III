package router

import (
	"database/sql"

	handlerodontologos "github.com/Damian-Damonte/practica-final-esp-back-III/cmd/server/handler/odontologos"
	"github.com/Damian-Damonte/practica-final-esp-back-III/cmd/server/handler/ping"
	"github.com/Damian-Damonte/practica-final-esp-back-III/internal/odontologos"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

func NewRouter(engine *gin.Engine, db *sql.DB) Router{
	return &router{
		engine: engine,
		db: db,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildPingRoutes()
	r.buildOdontologoRoutes()
}

func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

func (r *router) buildPingRoutes() {
	pingController := ping.NewControllerPing()
	r.routerGroup.GET("/ping", pingController.HandlerPing())
}

func (r *router) buildOdontologoRoutes() {
	repository := odontologos.NewMySqlRepository(r.db)
	service := odontologos.NewServiceOdontologo(repository)
	controlador := handlerodontologos.NewControladorOdontologo(service)

	grupoOdontologos := r.routerGroup.Group("/odontologos")
	{
		grupoOdontologos.GET("", controlador.HandlerGetAll())
		grupoOdontologos.GET(":id", controlador.HandlerGetById())
		grupoOdontologos.POST("", controlador.HandlerCreate())
		grupoOdontologos.PUT(":id", controlador.HandlerUpdate())
	}
}



