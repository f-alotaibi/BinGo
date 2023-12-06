package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetRawPasteContent(endpoint string) (string, error) {
	res, err := http.Get(endpoint)
	if err != nil || res.StatusCode != 200 {
		return "", fmt.Errorf("Could not get paste content")
	}
	body, berr := io.ReadAll(res.Body)
	if berr != nil {
		return "", fmt.Errorf("Could not get paste content")
	}
	return string(body), nil
}
