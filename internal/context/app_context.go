package context

import (
	"github.com/labstack/echo/v4"
	"github.com/ngonghi/admin_site/config"
	"github.com/ngonghi/admin_site/internal/cache"
	"github.com/ngonghi/admin_site/internal/repositories"
)

type AppContext struct {
	echo.Context
	Config *config.Configuration
	Cache  cache.Cache

	AdminRepository repositories.Admin
}
