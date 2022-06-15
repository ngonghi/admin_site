package middleware

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckAuth() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := session.Get("session", c)
			if err != nil {
				return c.String(http.StatusInternalServerError, "Error")
			}
			if b, _ := sess.Values["auth"]; b == nil {
				return c.Redirect(http.StatusMovedPermanently, "/login")
			}
			return h(c)
		}
	}
}
