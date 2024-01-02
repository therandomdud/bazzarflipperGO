package DataRequest

import (
	"io"
	"net/http"
)

// Defines the download sources
type URL string

// Difine here all URLs you need to download from.
const (
	BazaarAPI URL = "https://api.hypixel.net/skyblock/bazaar"
)

// GetRequest sends a GET request to the specified URL and returns the response body as a string
func (provider URL) GetRequest() (string, error) {
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
