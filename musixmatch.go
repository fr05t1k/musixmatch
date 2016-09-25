package musixmatch

import (
	"encoding/json"
	"fmt"
	"github.com/fr05t1k/musixmatch/config"
	"github.com/fr05t1k/musixmatch/entity"
	"github.com/fr05t1k/musixmatch/http"
	"github.com/fr05t1k/musixmatch/http/methods"
	"net/url"
)

type musixMatch struct {
	ApiKey string
}

func New(apiKey string) musixMatch {
	return musixMatch{ApiKey: apiKey}
}

func (m *musixMatch) GetLyrics(trackId uint32) (*entity.Lyrics, error) {
	var lyricsResponse entity.GetLyricsResponse
	params := url.Values{}
	params.Add(config.ApiKey, m.ApiKey)
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))

	resp, err := http.SendRequest(http.GetURLString(methods.TrackLyricsGet, params))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp, &lyricsResponse)

	if err != nil {
		return nil, err
	}

	return &lyricsResponse.Message.Body.Lyrics, nil
}
