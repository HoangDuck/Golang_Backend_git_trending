package middleware

import (
	"backend_github_trending/model"
	"backend_github_trending/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{Claims: &model.JwtCustomClaims{}, SigningKey: []byte(security.SECRET_KEY)}
	return middleware.JWTWithConfig(config)
}
