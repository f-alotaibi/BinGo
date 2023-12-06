package app

import (
	"bingo/internal/api"
	"net/http"

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
	var pastebin api.Paste = api.Pastebin{}
	var githubgists api.Paste = api.GithubGists{}
	var gitlabsnippets api.Paste = api.GitlabSnippets{}
	var rentry api.Paste = api.Rentry{}
	pastebin.Init(e.Group("/pastebin"))
	githubgists.Init(e.Group("/github"))
	gitlabsnippets.Init(e.Group("/gitlab"))
	rentry.Init(e.Group("/rentry"))
	e.Renderer = echoview.Default()
	return e
}
