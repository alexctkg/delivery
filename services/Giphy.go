package services

import (
	"encoding/json"
	"net/http"
)

// service para request na api Giphy
func GiphyRequest(title string) (*int, map[string]interface{}, error) {
	url := "http://api.giphy.com/v1/gifs/search?q=cheeseburgers&api_key=GQ2lrWPMUgLVvxNLpgCUk4PMmaZVtEW7"
	request, _ := http.NewRequest("POST", url, nil)
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
