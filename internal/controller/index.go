package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Index struct {
}

func (ctr Index) GetIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
