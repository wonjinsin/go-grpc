package router

import (
	ct "phantom/controller"
	"phantom/service"

	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo, svc *service.Service) {
	api := e.Group("/api")
	ver := api.Group("/v1")

	makeV1Route(ver, svc)
}

func makeV1Route(ver *echo.Group, svc *service.Service) {
	makeV1UserRoute(ver, svc)
}

func makeV1UserRoute(ver *echo.Group, svc *service.Service) {
	user := ver.Group("/user")
	userCt := ct.NewUserController(svc.User)
	user.GET("/:uid", userCt.GetUser)
	user.PUT("/:uid", userCt.UpdateUser)
	user.DELETE("/:uid", userCt.DeleteUser)
}
