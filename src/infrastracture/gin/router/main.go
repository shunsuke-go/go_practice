package gin

import (
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
	router.InitialRoute()
	router.Run(":3000")
}
