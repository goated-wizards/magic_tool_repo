package image

import (
	"errors"
	"path"
	"strings"
)

func GetExtension(url string) (string, error) {
	// strip query parameters
	noQuery := strings.Split(url, "?")[0]

	ext := strings.ToLower(path.Ext(noQuery)) // gets ".jpg", ".png", etc.

	switch ext {
	case ".png", ".jpg", ".jpeg":
		return ext, nil
	}
	return "", errors.New("format not found")

}
