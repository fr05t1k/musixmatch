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
	mm := musixMatch{ApiKey: getApiKey()}

	expectedId := uint32(113303287)
	response, err := mm.getLyrics(expectedId)

	assert.Nil(t, err)
	assert.Equal(t, uint16(200), response.Message.Header.StatusCode)
	assert.True(t, len(response.Message.Body.Lyrics.Body) > 0)
}
