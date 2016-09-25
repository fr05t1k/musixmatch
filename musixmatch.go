package musixmatch

import (
	"encoding/json"
	"fmt"
	"github.com/fr05t1k/musixmatch/config"
	"github.com/fr05t1k/musixmatch/entity/lyrics"
	"github.com/fr05t1k/musixmatch/entity/snippet"
	"github.com/fr05t1k/musixmatch/entity/subtitle"
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

// Get the lyrics of a track.
func (m *musixMatch) GetLyrics(trackId uint32) (*lyrics.Lyrics, error) {
	params := m.paramsWithApiKey()
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))

	resp, err := http.SendRequest(http.GetURLString(methods.TrackLyricsGet, params))
	if err != nil {
		return nil, err
	}
	var lyricsResponse lyrics.Response
	err = json.Unmarshal(resp, &lyricsResponse)

	if err != nil {
		return nil, err
	}

	return &lyricsResponse.Message.Body.Lyrics, nil
}

// Get the snippet for a given track.
//
// A lyrics snippet is a very short representation of a song lyrics.
// It’s usually twenty to a hundred characters long and it’s calculated
// extracting a sequence of words from the lyrics.
func (m *musixMatch) GetSnippet(trackId uint32) (*snippet.Snippet, error) {

	params := m.paramsWithApiKey()
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))
	resp, err := http.SendRequest(http.GetURLString(methods.TrackSnippetGet, params))

	var snippetResponse snippet.Response
	err = json.Unmarshal(resp, &snippetResponse)

	if err != nil {
		return nil, err
	}

	return &snippetResponse.Message.Body.Snippet, nil
}

// Retrieve the subtitle of a track.
//
// Return the subtitle of a track in LRC or DFXP format.
func (m *musixMatch) GetSubtitles(trackId uint32) (*subtitle.Subtitle, error) {

	params := m.paramsWithApiKey()
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))
	resp, err := http.SendRequest(http.GetURLString(methods.TrackSubtitleGet, params))

	var subtitleResponse subtitle.Response
	err = json.Unmarshal(resp, &subtitleResponse)

	if err != nil {
		return nil, err
	}

	return &subtitleResponse.Message.Body.Subtitle, nil
}

// Retrieve the url.Values with predefined api key
func (m *musixMatch) paramsWithApiKey() url.Values {
	params := url.Values{}
	params.Add(config.ApiKey, m.ApiKey)
	return params
}
