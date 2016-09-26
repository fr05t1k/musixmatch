package musixmatch

import (
	"github.com/fr05t1k/musixmatch/http"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getApiKey() string {
	key := os.Getenv("MUSIXMATCH_API_KEY")
	if key == "" {
		panic("Please specify MUSIXMATCH_API_KEY environment variable")
	}
	return key
}

func TestGetLyrics(t *testing.T) {

	mm := New(getApiKey())

	expectedId := uint32(113303287)
	lyrics, err := mm.GetLyrics(expectedId)

	assert.Nil(t, err)
	assert.True(t, len(lyrics.Body) > 0)
}

func TestMusixMatch_GetSnippet(t *testing.T) {
	mm := New(getApiKey())

	expectedId := uint32(113303287)
	snippet, err := mm.GetSnippet(expectedId)

	assert.Nil(t, err)
	assert.True(t, len(snippet.Body) > 0)
}

func TestMusixMatch_GetSubtitles(t *testing.T) {
	mm := New(getApiKey())

	expectedId := uint32(113303287)
	subtitle, _ := mm.GetSubtitles(expectedId)

	assert.Nil(t, subtitle)
}

func TestMusixMatch_SearchTrack(t *testing.T) {
	mm := New(getApiKey())

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
