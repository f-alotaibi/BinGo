package api

import (
	"net/http"
	"pastebin-go/internal/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetDownloadPaste(c echo.Context) error {
	content, err := utils.GetPasteContent(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	r := strings.NewReader(content)
	return c.Stream(http.StatusOK, "application/octet-stream", r)
}
