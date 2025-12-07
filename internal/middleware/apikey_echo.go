package middleware

import (
	"net/http"

	"github.com/Mayer-04/go-mongo-people-courses-api/internal/config"
	"github.com/labstack/echo/v4"
)

type APIKeyErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func APIKeyMiddleware(cfg config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("X-API-KEY")

			if key == "" || key != cfg.ApiKey {
				return c.JSON(http.StatusForbidden, APIKeyErrorResponse{
					Error:   "forbidden",
					Message: "invalid or missing API key",
				})
			}

			return next(c)
		}
	}
}
