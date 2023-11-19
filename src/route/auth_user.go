package route

import (
	"github.com/kouhei-github/fiber-sample-framework/controller"
	"github.com/kouhei-github/fiber-sample-framework/middlewares"
)

func (router *Router) GetAuthRouter() {
	router.FiberApp.Post("/api/v1/signup", controller.SignUpHandler)
	router.FiberApp.Post("/api/v1/login", controller.LoginHandler)
	router.FiberApp.Get("/api/v1/user", middlewares.CheckJwtToken, controller.GetAllUserHandler)
}
