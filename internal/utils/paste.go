package utils

import (
	"fmt"
	"io"
	"net/http"
)

var ENDPOINT = "https://pastebin.com/raw/%s"

func GetPasteContent(id string) (string, error) {
	res, err := http.Get(fmt.Sprintf(ENDPOINT, id))
	if err != nil || res.StatusCode != 200 {
		return "", fmt.Errorf("Could not get paste content")
	}
	body, berr := io.ReadAll(res.Body)
	if berr != nil {
		return "", fmt.Errorf("Could not get paste content")
	}
	return string(body), nil
}
