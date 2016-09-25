package http

import (
	"github.com/fr05t1k/musixmatch/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
