package core

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	var errorPage string

	switch code {
	case 400, 403, 404, 500:
		errorPage = fmt.Sprintf("errors/%d.html", code)
	default:
		errorPage = "errors/500.html"
	}
	
	c.Logger().Error(errorPage)

	if err := c.Render(code, errorPage, nil); err != nil {
		c.Logger().Error(err)
	}
}
