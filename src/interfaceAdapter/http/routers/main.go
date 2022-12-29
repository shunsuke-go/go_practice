package routers

import (
	"go_practice/infrastracture/db"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func createRouter() *Router {
	return &Router{
		gin.Default(),
	}
}

func Run() {
	router := createRouter()

	postgresConnector := db.NewPostgresConnector()
	router.InitialRoute(postgresConnector.Conn)
	router.Run(":3000")
}
