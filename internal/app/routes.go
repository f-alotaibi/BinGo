package app

import (
	"net/http"
	"pastebin-go/internal/api"

	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func loadRoutes() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.GET("/id/:id", api.GetPasteID)
	e.GET("/idnew/:id", api.GetPasteNewID)
	e.GET("/raw/:id", api.GetPasteRaw)
	e.GET("/dl/:id", api.GetDownloadPaste)

	e.Renderer = echoview.Default()
	return e
}
