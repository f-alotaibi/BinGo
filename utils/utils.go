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

func FormatStringSize(size float64) string {
	var suffixes = []string{"B", "KB", "MB", "GB", "TB"}
	suffixIndex := 0
	for size > 1024 && suffixIndex < 4 {
		suffixIndex++
		size /= 1024
	}
	if suffixIndex == 0 {
		return fmt.Sprintf("%.0f %s", size, suffixes[suffixIndex])
	}
	return fmt.Sprintf("%.2f %s", size, suffixes[suffixIndex])
}
