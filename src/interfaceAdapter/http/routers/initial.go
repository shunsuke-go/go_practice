package routers

import (
	"database/sql"
	"go_practice/interfaceAdapter/http/handler"

	"github.com/gin-gonic/gin"
)

func (router *Router) InitialRoute(db *sql.DB) {
	router.GET("/", func (c *gin.Context) {
		handler.Inital(c, db)
	})
}
