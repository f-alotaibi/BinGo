package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type App struct {
	router *echo.Echo
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Start() error {
	err := a.router.Start(":80")
	if err != nil {
		a.router.StdLogger.Fatal(fmt.Errorf("Failed to start server: %w", err))
		return err
	}
	return nil
}
