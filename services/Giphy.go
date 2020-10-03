package services

import (
	"encoding/json"
	"net/http"
	"os"
)

//GiphyRequest service para request na api Giphy
func GiphyRequest(title string) (*int, map[string]interface{}, error) {
	url := "https://api.giphy.com/v1/gifs/search?limit=1&api_key=" + os.Getenv("GIPHY_API_KEY") + "&q=?" + title
	request, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	decoder := json.NewDecoder(response.Body)
	var ret map[string]interface{}

	err = decoder.Decode(&ret)
	if err != nil {

		return nil, nil, err
	}

	return &response.StatusCode, ret, nil
}
