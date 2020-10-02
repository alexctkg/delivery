package services

import (
	"encoding/json"
	"net/http"
)

func RecipePuppyRequest(ingredientes string) (*int, map[string]interface{}, error) {
	url := "http://www.recipepuppy.com/api/?i=" + ingredientes
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
