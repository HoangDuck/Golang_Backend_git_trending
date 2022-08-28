package middleware

import (
	"backend_github_trending/model"
	req2 "backend_github_trending/model/req"
	"github.com/labstack/echo"
	"net/http"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {
			//handle logic
			req := req2.ReqSignIn{}
			if err := context.Bind(&req); err != nil {
				return context.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			if req.Email != "admin@gmail.com" {
				return context.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "Bạn không có quyền",
					Data:       nil,
				})
			}
			return next(context)
		}
	}
}
