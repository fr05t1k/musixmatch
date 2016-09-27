/*
Package musixmatch provides a client for using the Musixmatch API.
*/
package musixmatch

import (
	"fmt"
	"github.com/fr05t1k/musixmatch/config"
	"github.com/fr05t1k/musixmatch/entity/lyrics"
	"github.com/fr05t1k/musixmatch/entity/snippet"
	"github.com/fr05t1k/musixmatch/entity/subtitle"
	"github.com/fr05t1k/musixmatch/entity/track"
	"github.com/fr05t1k/musixmatch/entity/track/list"
	"github.com/fr05t1k/musixmatch/http"
	"github.com/fr05t1k/musixmatch/http/methods"
	"net/url"
)

type Client struct {
	ApiKey string
}

func NewClient(apiKey string) Client {
	return Client{ApiKey: apiKey}
}

// Get the lyrics of a track.
func (m *Client) GetLyrics(trackId uint32) (*lyrics.Lyrics, error) {
	params := m.paramsWithApiKey()
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))

	var lyricsResponse lyrics.Response
	requestUrl := http.GetURLString(methods.TrackLyricsGet, params)
	err := http.RequestEntity(requestUrl, &lyricsResponse)

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
func (m *Client) GetSnippet(trackId uint32) (*snippet.Snippet, error) {

	params := m.paramsWithApiKey()
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))
	requestUrl := http.GetURLString(methods.TrackSnippetGet, params)
	var snippetResponse snippet.Response
	err := http.RequestEntity(requestUrl, &snippetResponse)

	if err != nil {
		return nil, err
	}

	return &snippetResponse.Message.Body.Snippet, nil
}

// Retrieve the subtitle of a track.
//
// Return the subtitle of a track in LRC or DFXP format.
func (m *Client) GetSubtitles(trackId uint32) (*subtitle.Subtitle, error) {

	params := m.paramsWithApiKey()
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))

	var subtitleResponse subtitle.Response
	requestUrl := http.GetURLString(methods.TrackSubtitleGet, params)
	err := http.RequestEntity(requestUrl, &subtitleResponse)
	if err != nil {
		return nil, err
	}

	return &subtitleResponse.Message.Body.Subtitle, nil
}

// Search tracks using the given criteria.
func (m *Client) SearchTrack(request http.SearchRequest) ([]list.TrackList, error) {
	params := m.paramsWithApiKey()
	params.Add(config.Query, request.Query)
	params.Add(config.QueryArtist, request.QueryArtist)
	params.Add(config.QueryTrack, request.QueryTrack)
	params.Add(config.Page, fmt.Sprintf("%d", request.Page))
	params.Add(config.PageSize, fmt.Sprintf("%d", request.PageSize))
	if request.HasLyrics {
		params.Add(config.FHasLyrics, "1")
	} else {
		params.Add(config.FHasLyrics, "0")
	}

	var tracksResponse list.Response
	requestUrl := http.GetURLString(methods.TrackSearch, params)
	err := http.RequestEntity(requestUrl, &tracksResponse)

	if err != nil {
		return []list.TrackList{}, err
	}

	return tracksResponse.Message.Body.Tracks, nil
}

// Get a track info from a database: title, artist, instrumental flag and cover art.
func (m *Client) GetTrack(id uint32) (*track.Track, error) {

	params := m.paramsWithApiKey()
	params.Add(config.TrackId, fmt.Sprintf("%d", id))

	var trackResponse track.Response
	requestUrl := http.GetURLString(methods.TrackGet, params)
	err := http.RequestEntity(requestUrl, &trackResponse)
	if err != nil {
		return nil, err
	}

	return &trackResponse.Message.Body.Track, nil
}

// Match your song against a database.
func (m *Client) GetMatchingTrack(qTrack string, qArtist string) (*track.Track, error) {

	params := m.paramsWithApiKey()
	params.Add(config.QueryTrack, qTrack)
	params.Add(config.QueryArtist, qArtist)

	var trackResponse track.Response
	requestUrl := http.GetURLString(methods.MatcherTrackGet, params)
	err := http.RequestEntity(requestUrl, &trackResponse)
	if err != nil {
		return nil, err
	}

	return &trackResponse.Message.Body.Track, nil
}

// Make request with your parameters
func (m *Client) Request(method string, params url.Values) ([]byte, error) {
	params.Add(config.ApiKey, m.ApiKey)

	requestUrl := http.GetURLString(method, params)
	return http.SendRequest(requestUrl)
}

// Retrieve the url.Values with predefined api key
func (m *Client) paramsWithApiKey() url.Values {
	params := url.Values{}
	params.Add(config.ApiKey, m.ApiKey)
	return params
}
