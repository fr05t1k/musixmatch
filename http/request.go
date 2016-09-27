package http

import (
	"encoding/json"
	"github.com/fr05t1k/musixmatch/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SearchRequest struct {
	Query       string
	QueryArtist string
	QueryTrack  string
	Page        uint32
	PageSize    uint8
	HasLyrics   bool
}

func SendRequest(request string) ([]byte, error) {
	apiUrl := config.ApiUrl + config.ApiVersion + config.UrlDelim + request
	response, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func GetURLString(method string, params url.Values) string {
	return method + "?" + params.Encode()
}

// Retrieve the entity
func RequestEntity(url string, entity interface{}) error {
	resp, err := SendRequest(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, entity)
	return err
}
