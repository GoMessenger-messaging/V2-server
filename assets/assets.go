package assets

import (
	"net/http"
	"os"
	"strings"
)

func GetAsset(location string, dir string) (data []byte, contentType string, err error) {
	path := strings.TrimSuffix(dir, "/") + "/" + location
	return getFile(path)
}

func getFile(path string) ([]byte, string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, "", err
	}
	contentType := http.DetectContentType(data)
	return data, contentType, nil
}
