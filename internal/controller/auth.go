package controller

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/ngonghi/admin_site/internal/core"
	"net/http"
)

type (
	Auth struct {
	}

	LoginForm struct {
		Email    string `form:"email" transFieldName:"メール" validate:"required,email"`
		Password string `form:"password" transFieldName:"パスワード" validate:"required"`
	}

	LoginViewModel struct {
		CSRF              string
		EmailErrorMess    string
		PasswordErrorMess string
	}
)

func (ctr Auth) GetLogin(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error")
	}

	vm := LoginViewModel{
		CSRF: c.Get("csrf").(string),
	}

	if sess.Values["errors"] != nil {
		errors := sess.Values["errors"].(map[string]string)
		if errors["LoginForm.パスワード"] != "" {
			vm.PasswordErrorMess = errors["LoginForm.パスワード"]
		}
		if errors["LoginForm.メール"] != "" {
			vm.EmailErrorMess = errors["LoginForm.メール"]
		}
	}

	sess.Values["errors"] = nil
	sess.Save(c.Request(), c.Response())

	return c.Render(http.StatusOK, "auth/login.html", vm)
}

func (ctr Auth) PostLogin(c echo.Context) (err error) {
	sess, _ := session.Get("session", c)

	form := new(LoginForm)
	if err = c.Bind(form); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(form); err != nil {
		sess.Values["errors"] = core.GetErrorMessages(err)
		if e := sess.Save(c.Request(), c.Response()); e != nil {
			c.Logger().Error(e)
		}
		return c.Redirect(http.StatusMovedPermanently, "/login")
	}
	
	//id := c.FormValue("id")
	//password := c.FormValue("password")

	sess.Values["auth"] = "vian"
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}
