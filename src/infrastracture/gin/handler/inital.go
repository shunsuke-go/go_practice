package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go_practice/interfaceAdapter/http/presenter"
)

func Inital(c *gin.Context) {
	// コントローラーをインポートしてここで使う
	c.JSON(http.StatusOK, presenter.Initial())
}
