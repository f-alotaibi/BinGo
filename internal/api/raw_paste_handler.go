package api

import (
	"net/http"
	"pastebin-go/internal/utils"

	"github.com/labstack/echo/v4"
)

func GetPasteRaw(c echo.Context) error {
	content, err := utils.GetPasteContent(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Not Found")
	}
	return c.Blob(http.StatusOK, "text/plain", []byte(content))
}
