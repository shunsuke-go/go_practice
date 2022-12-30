package routers

import (
	"go_practice/infrastracture/db"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	return &Router{
		gin.Default(),
	}
}

func Run() {
	router := NewRouter()

	postgresConnector := db.NewPostgresConnector()
	router.InitialRoute(postgresConnector.Conn)
	router.Run(":3000")
}
