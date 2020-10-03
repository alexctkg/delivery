package models

import "encoding/json"

// GiphyResponse struct de retorno da api giphy
type GiphyResponse struct {
	Data []GiphyData `json:"data"`
}

//GiphyData .
type GiphyData struct {
	EmbedURL string `json:"embed_url"`
	ID       string `json:"id"`
}

//MapToGiphyURL extração da url do retorno da api
func MapToGiphyURL(mapToConvert map[string]interface{}) (*string, error) {
	var url string
	giphy := GiphyResponse{}

	jsonbody, err := json.Marshal(mapToConvert)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(jsonbody, &giphy); err != nil {
		return nil, err
	}

	if len(giphy.Data) > 0 { //confere se exite o array
		//extrai a EmbedURL, poderia ser qualquer outro gif do retorno da api, basta tratar aqui
		url = giphy.Data[0].EmbedURL
	}

	return &url, nil

}
