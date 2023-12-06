package api

import (
	"bingo/internal/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Rentry struct{}

func (paste Rentry) Init(g *echo.Group) {
	g.GET("/id/:id", paste.getPasteID)
	g.GET("/raw/:id", paste.getPasteRaw)
	g.GET("/dl/:id", paste.getDownloadPaste)
}

func (paste Rentry) getEndpoint() string {
	return "https://rentry.co/%s/raw"
}

func (paste Rentry) getPastename() string {
	return "rentry"
}

func (paste Rentry) getPasteID(c echo.Context) error {
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

func (paste Rentry) getPasteRaw(c echo.Context) error {
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	return c.Blob(http.StatusOK, "text/plain", []byte(content))
}

func (paste Rentry) getDownloadPaste(c echo.Context) error {
	content, err := utils.GetRawPasteContent(fmt.Sprintf(paste.getEndpoint(), c.Param("id")))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	r := strings.NewReader(content)
	return c.Stream(http.StatusOK, "application/octet-stream", r)
}
