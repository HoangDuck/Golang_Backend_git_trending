package router

import (
	"backend_github_trending/handler"
	middleware "backend_github_trending/middleware"
	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetUpRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
	api.Echo.POST("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
}
