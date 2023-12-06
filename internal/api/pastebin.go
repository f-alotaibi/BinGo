package api

import (
	"fmt"
	"net/http"
	"pastebin-go/internal/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

type Pastebin struct{}

func (paste Pastebin) Init(g *echo.Group) {
	g.GET("/id/:id", paste.getPasteID)
	g.GET("/raw/:id", paste.getPasteRaw)
	g.GET("/dl/:id", paste.getDownloadPaste)
}

func (paste Pastebin) getEndpoint() string {
	return "https://pastebin.com/raw/%s"
}

func (paste Pastebin) getPastename() string {
	return "pastebin"
}

func (paste Pastebin) getPasteID(c echo.Context) error {
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

func (paste Pastebin) getPasteRaw(c echo.Context) error {
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	return c.Blob(http.StatusOK, "text/plain", []byte(content))
}

func (paste Pastebin) getDownloadPaste(c echo.Context) error {
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	r := strings.NewReader(content)
	return c.Stream(http.StatusOK, "application/octet-stream", r)
}
