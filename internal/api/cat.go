package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CatImage() (string, error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/images/search")
	if err != nil {
		return "", fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK status code: %v", resp.StatusCode)
	}

	var cats []struct {
		URL string `json:"url"`
	}
	err = json.NewDecoder(resp.Body).Decode(&cats)
	if err != nil {
		return "", fmt.Errorf("failed to decode response body: %v", err)
	}

	if len(cats) == 0 {
		return "", fmt.Errorf("no cat image URL found in response")
	}

	return cats[0].URL, nil
}
