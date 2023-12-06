package api

import (
	"bingo/internal/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type GitlabSnippets struct{}

func (paste GitlabSnippets) Init(g *echo.Group) {
	g.GET("/id/:id", paste.getPasteID)
	g.GET("/raw/:id", paste.getPasteRaw)
	g.GET("/dl/:id", paste.getDownloadPaste)
}

func (paste GitlabSnippets) getEndpoint() string {
	return "https://gitlab.com/snippets/%s/raw"
}

func (paste GitlabSnippets) getPastename() string {
	return "gitlab"
}

func (paste GitlabSnippets) getPasteID(c echo.Context) error {
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

func (paste GitlabSnippets) getPasteRaw(c echo.Context) error {
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	return c.Blob(http.StatusOK, "text/plain", []byte(content))
}

func (paste GitlabSnippets) getDownloadPaste(c echo.Context) error {
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	r := strings.NewReader(content)
	return c.Stream(http.StatusOK, "application/octet-stream", r)
}
