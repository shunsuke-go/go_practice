package gin

import "go_practice/infrastracture/gin/handler"

func (router *Router) InitialRoute() {
	router.GET("/", handler.Inital)
}
