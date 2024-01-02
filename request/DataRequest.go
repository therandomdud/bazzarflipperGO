package DataRequest

import (
	"io"
	"net/http"
)

// Defines the download sources
type URLProvider string

// Difine here all URLs you need to download from.
const (
	BazaarAPI URLProvider = "https://api.hypixel.net/skyblock/bazaar"
)

// GetRequest sends a GET request to the specified URL and returns the response body as a string
func (provider URLProvider) GetRequest() (string, error) {
	client := http.Client{}
	request, err := http.NewRequest("GET", string(provider), nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
