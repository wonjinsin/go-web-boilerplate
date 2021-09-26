package router

import (
	ct "pikachu/controller"
	"pikachu/repository"
	"pikachu/service"

	"github.com/labstack/echo"
)

// Init ...
func Init(e *echo.Echo, svc *service.Service, repo *repository.Repository) {
	// Default Group
	api := e.Group("/api")
	ver := api.Group("/v1")

	// User Controller
	user := ver.Group("/user")
	userCt := ct.NewUserController(svc.User)
	user.POST("", userCt.NewUser)
	user.GET("/:uid", userCt.GetUser)
	user.PUT("/:uid", userCt.UpdateUser)
	user.DELETE("/:uid", userCt.DeleteUser)
}
