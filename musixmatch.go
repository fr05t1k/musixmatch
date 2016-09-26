package musixmatch

import (
	"fmt"
	"github.com/fr05t1k/musixmatch/config"
	"github.com/fr05t1k/musixmatch/entity/lyrics"
	"github.com/fr05t1k/musixmatch/entity/snippet"
	"github.com/fr05t1k/musixmatch/entity/subtitle"
	"github.com/fr05t1k/musixmatch/entity/track"
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
func (m *musixMatch) GetSnippet(trackId uint32) (*snippet.Snippet, error) {

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
func (m *musixMatch) GetSubtitles(trackId uint32) (*subtitle.Subtitle, error) {

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
func (m *musixMatch) SearchTrack(request http.SearchRequest) ([]track.TrackList, error) {
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

	var tracksResponse track.Response
	requestUrl := http.GetURLString(methods.TrackSearch, params)
	err := http.RequestEntity(requestUrl, &tracksResponse)

	if err != nil {
		return []track.TrackList{}, err
	}

	return tracksResponse.Message.Body.Tracks, nil
}

// Retrieve the url.Values with predefined api key
func (m *musixMatch) paramsWithApiKey() url.Values {
	params := url.Values{}
	params.Add(config.ApiKey, m.ApiKey)
	return params
}
