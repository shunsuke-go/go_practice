package routers

import (
	"database/sql"

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

func Run(db *sql.DB) {
	router := NewRouter()

	router.InitialRoute(db)
	router.Run(":3000")
}
