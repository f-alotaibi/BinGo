package api

import "github.com/labstack/echo/v4"

type Paste interface {
	Init(g *echo.Group)
	getEndpoint() string
	getPastename() string
	getPasteID(c echo.Context) error
	getPasteRaw(c echo.Context) error
	getDownloadPaste(c echo.Context) error
}
