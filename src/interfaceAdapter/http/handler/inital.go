package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"go_practice/interfaceAdapter/gateway"
	"go_practice/interfaceAdapter/http/presenter"
	"go_practice/usecase/interactor"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Inital(c *gin.Context, db *sql.DB) {
	initialRepository := gateway.NewInitialRepository(db)
	initialInteractor := interactor.NewInitialInteractor(initialRepository)

	user, err := initialInteractor.FindOneUser(c)
	if err != nil {
		fmt.Print("db error", err)
	}

	initialPresenter := presenter.NewInitialPresenter(user)
	c.JSON(http.StatusOK, initialPresenter.ToInitialPresenter())
}
