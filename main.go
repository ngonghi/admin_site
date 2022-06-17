package main

import (
	"encoding/gob"
	"github.com/ngonghi/admin_site/config"
	"github.com/ngonghi/admin_site/internal/controller"
	"github.com/ngonghi/admin_site/internal/core"
	"github.com/ngonghi/admin_site/internal/middleware"
	"github.com/ngonghi/admin_site/internal/models"
	"log"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	gob.Register(map[string]string{})

	// create server
	server := core.NewServer(config)
	// serve files for dev
	server.ServeStaticFiles()

	indexCtrl := &controller.Index{}
	authCtrl := &controller.Auth{}

	// pages
	server.Echo.GET("", indexCtrl.GetIndex, middleware.CheckAuth())

	server.Echo.GET("/login", authCtrl.GetLogin)
	server.Echo.POST("/login", authCtrl.PostLogin)

	// migration for dev
	admin := models.Admin{}
	mr := server.GetModelRegistry()
	err = mr.Register(admin)

	if err != nil {
		server.Echo.Logger.Fatal(err)
	}

	mr.AutoMigrateAll()

	// Start server
	go func() {
		if err := server.Start(config.Address); err != nil {
			server.Echo.Logger.Info("shutting down the server")
		}
	}()

	server.GracefulShutdown()
}
