package api

import (
	"fmt"
	"net/http"
	"pastebin-go/internal/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

type GithubGists struct{}

func (paste GithubGists) Init(g *echo.Group) {
	g.GET("/id/:id", paste.getPasteID)
	g.GET("/raw/:id", paste.getPasteRaw)
	g.GET("/dl/:id", paste.getDownloadPaste)
}

func (paste GithubGists) getEndpoint() string {
	return "https://gist.githubusercontent.com/%s/raw"
}

func (paste GithubGists) getPastename() string {
	return "github"
}

func (paste GithubGists) getPasteID(c echo.Context) error {
	id := c.Param("id")
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	return c.Render(http.StatusOK, "paste.html", map[string]interface{}{
		"id":      id,
		"bin":     paste.getPastename(),
		"content": content,
	})
}

func (paste GithubGists) getPasteRaw(c echo.Context) error {
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	return c.Blob(http.StatusOK, "text/plain", []byte(content))
}

func (paste GithubGists) getDownloadPaste(c echo.Context) error {
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	r := strings.NewReader(content)
	return c.Stream(http.StatusOK, "application/octet-stream", r)
}
