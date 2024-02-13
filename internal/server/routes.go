package server

import (
	"bingo/api"
	"bingo/web"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := httprouter.New()
	r.Handler(http.MethodGet, "/", templ.Handler(web.HomeComponent()))

	api.Pastebin.Init(r)
	api.GithubGists.Init(r)
	api.GitlabSnippets.Init(r)
	api.RentryCo.Init(r)
	return r
}
