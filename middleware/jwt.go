package middleware

import (
	"backend_github_trending/model"
	"backend_github_trending/security"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{Claims: &model.JwtCustomClaims{}, SigningKey: security.SECRET_KEY}
	return middleware.JWTWithConfig(config)
}
