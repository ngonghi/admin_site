package controller

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Auth struct {
	}

	LoginViewModel struct {
		CSRF string
	}
)

func (ctr Auth) GetLogin(c echo.Context) error {
	c.Logger().Error(c.Get("csrf"))

	vm := LoginViewModel{
		CSRF: c.Get("csrf").(string),
	}

	return c.Render(http.StatusOK, "auth/login.html", vm)
}

func (ctr Auth) PostLogin(c echo.Context) error {
	id := c.FormValue("id")
	password := c.FormValue("password")

	c.Logger().Error(id)
	c.Logger().Error(password)

	sess, _ := session.Get("session", c)

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values["auth"] = "vian"
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}
