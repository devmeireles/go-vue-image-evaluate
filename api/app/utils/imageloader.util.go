package utils

import (
	"bytes"
	"io"
	"net/http"
)

func ImageLoader(url string) (*bytes.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	imgBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(imgBytes), nil
}
