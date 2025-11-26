package image

import (
	"io"
	"net/http"
	"os"
)

func DownloadImage(url, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy data from HTTP response to the file
	_, err = io.Copy(out, resp.Body)
	return err
}
