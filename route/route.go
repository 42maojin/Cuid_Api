package route

import (
	"project_api/handler"

	"github.com/julienschmidt/httprouter"
)

// NewRouter 路由配置
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", handler.InsertNode)
	router.DELETE("/user", handler.DeleteNode)
	router.PUT("/user", handler.UpdateNode)
	router.GET("/user", handler.SeleNode)
	return router
}
