package api

import (
	"net/http"
	"pastebin-go/internal/utils"

	"github.com/labstack/echo/v4"
)

func GetPasteID(c echo.Context) error {
	id := c.Param("id")
	content, err := utils.GetPasteContent(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	return c.Render(http.StatusOK, "paste.html", map[string]interface{}{
		"id":      id,
		"content": content,
	})
}

func GetPasteNewID(c echo.Context) error {
	id := c.Param("id")
	content, err := utils.GetPasteContent(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	return c.Render(http.StatusOK, "pastenew.html", map[string]interface{}{
		"id":      id,
		"content": content,
	})
}
