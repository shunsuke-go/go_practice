package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shunsuke-go/go_practice/interfaceAdapter/http/presenter"
)

func Inital(c *gin.Context) {
	// コントローラーをインポートしてここで使う
	c.JSON(http.StatusOK, presenter.Initial())
}
