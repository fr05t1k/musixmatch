package musixmatch

import (
	"encoding/json"
	"github.com/fr05t1k/musixmatch/config"
	"github.com/fr05t1k/musixmatch/http"
	"github.com/fr05t1k/musixmatch/http/methods"
	"github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
	//"github.com/fr05t1k/musixmatch/entity/track"
	"fmt"
	"github.com/fr05t1k/musixmatch/entity/track"
)

const trackId uint32 = 113303287

func getApiKey() string {
	key := os.Getenv("MUSIXMATCH_API_KEY")
	if key == "" {
		panic("Please specify MUSIXMATCH_API_KEY environment variable")
	}
	return key
}

func TestGetLyrics(t *testing.T) {

	mm := NewClient(getApiKey())

	expectedId := uint32(trackId)
	lyrics, err := mm.GetLyrics(expectedId)

	assert.Nil(t, err)
	assert.True(t, len(lyrics.Body) > 0)
}

func TestMusixMatch_GetSnippet(t *testing.T) {
	mm := NewClient(getApiKey())

	expectedId := uint32(trackId)
	snippet, err := mm.GetSnippet(expectedId)

	assert.Nil(t, err)
	assert.True(t, len(snippet.Body) > 0)
}

func TestMusixMatch_GetSubtitles(t *testing.T) {
	mm := NewClient(getApiKey())

	expectedId := uint32(trackId)
	subtitle, _ := mm.GetSubtitles(expectedId)

	assert.Nil(t, subtitle)
}

func TestMusixMatch_SearchTrack(t *testing.T) {
	mm := NewClient(getApiKey())

	request := http.SearchRequest{
		Query: "Heathens",
	}
	tracks, err := mm.SearchTrack(request)

	assert.Nil(t, err)
	assert.NotEmpty(t, tracks)
	trackList := tracks[0]

	assert.NotEmpty(t, trackList.Track.Id)
	assert.NotEmpty(t, trackList.Track.AlbumId)
	assert.NotEmpty(t, trackList.Track.Name)
	assert.NotEmpty(t, trackList.Track.Rating)
}

func TestMusixMatch_GetTrack(t *testing.T) {
	mm := NewClient(getApiKey())
	trackEntity, err := mm.GetTrack(trackId)
	assert.Nil(t, err)
	assert.NotEmpty(t, trackEntity)
}

func TestMusixMatch_GetMatchingTrack(t *testing.T) {
	mm := NewClient(getApiKey())
	trackEntity, err := mm.GetMatchingTrack("Yesterday", "Beatles")
	assert.Nil(t, err)
	assert.NotEmpty(t, trackEntity)
}

func TestMusixMatch_Request(t *testing.T) {
	mm := NewClient(getApiKey())

	params := url.Values{}
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))
	resp, err := mm.Request(methods.TrackGet, params)

	assert.Nil(t, err)

	var trackResponse track.Response

	json.Unmarshal(resp, &trackResponse)

	assert.Equal(t, uint16(200), trackResponse.Message.Header.StatusCode)
}

// Basic usages
func ExampleNewClient() {
	client := NewClient("{apikey}")

	lyrics, err := client.GetLyrics(trackId)

	if err != nil {
		fmt.Println(lyrics.Body)
	}
}

// Search track request
func ExampleClient_SearchTrack() {
	mm := NewClient(getApiKey())

	request := http.SearchRequest{
		Query: "Heathens",
	}
	tracks, err := mm.SearchTrack(request)
	if err != nil {
		fmt.Println(len(tracks))
	}
}

func ExampleClient_Request() {
	mm := NewClient(getApiKey())

	// prepare a method and a query
	params := url.Values{}
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))
	resp, _ := mm.Request(methods.TrackGet, params)

	var trackResponse track.Response

	// Let's hydrate response
	json.Unmarshal(resp, &trackResponse)
}
