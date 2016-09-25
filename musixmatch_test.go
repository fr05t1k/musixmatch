package musixmatch

import (
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
