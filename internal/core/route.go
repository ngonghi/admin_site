package core

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ngonghi/admin_site/internal/cache"
	"github.com/ngonghi/admin_site/internal/context"
	mid "github.com/ngonghi/admin_site/internal/middleware"
	"github.com/ngonghi/admin_site/internal/repositories"
	"gopkg.in/go-playground/validator.v9"
)

func NewRouter(server *Server) *echo.Echo {
	config := server.config
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}

	cc := context.AppContext{
		Cache:           &cache.RedisCache{Client: server.cache},
		Config:          config,
		AdminRepository: &repositories.AdminRepository{DB: server.db},
	}

	e.Use(mid.AppContext(&cc))

	if config.RequestLogger {
		e.Use(middleware.Logger()) // request logger
	}

	e.Use(middleware.Recover())       // panic errors are thrown
	e.Use(middleware.BodyLimit("5M")) // limit body payload to 5MB
	e.Use(middleware.Secure())        // provide protection against injection attacks
	e.Use(middleware.RequestID())     // generate unique requestId

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf_token",
	}))
	
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(config.SessionSecret))))

	// Add html templates with go template syntax
	renderer := newTemplateRenderer(config.LayoutDir, config.TemplateDir)
	e.Renderer = renderer

	e.HTTPErrorHandler = HTTPErrorHandler

	return e
}
