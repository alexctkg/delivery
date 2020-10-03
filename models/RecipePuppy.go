package models

import (
	"encoding/json"
)

// RecipePuppyResponse struct de retorno da api RecipePuppy
type RecipePuppyResponse struct {
	Href    string               `json:"href"`
	Results []RecipePuppyResults `json:"results"`
	Title   string               `json:"title"`
	Version float64              `json:"version"`
}

type RecipePuppyResults struct {
	Href        string  `json:"href"`
	Ingredients string  `json:"ingredients"`
	Thumbnail   string  `json:"thumbnail"`
	Title       string  `json:"title"`
	Giphy       *string `json:"-"`
}

//MapToRecipePuppy entra com um map e retorna a struct para manipulação
func MapToRecipePuppy(mapToConvert map[string]interface{}) (RecipePuppyResponse, error) {
	recipePuppy := RecipePuppyResponse{}

	jsonbody, err := json.Marshal(mapToConvert)
	if err != nil {
		return recipePuppy, err
	}

	if err := json.Unmarshal(jsonbody, &recipePuppy); err != nil {
		return recipePuppy, err
	}

	return recipePuppy, nil

}
