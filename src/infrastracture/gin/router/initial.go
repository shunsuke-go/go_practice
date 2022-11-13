package gin

import "github.com/shunsuke-go/go_practice/infrastracture/gin/handler"

func (router *Router) InitialRoute() {
	router.GET("/", handler.Inital)
}
