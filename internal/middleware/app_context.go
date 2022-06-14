package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/ngonghi/admin_site/internal/context"
)

func AppContext(ac *context.AppContext) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac.Context = c
			return h(ac)
		}
	}
}
