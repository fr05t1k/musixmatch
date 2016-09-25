package musixmatch

import (
	"fmt"
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
	fmt.Println(lyrics.Id)
	fmt.Println(lyrics.Body)
	fmt.Println(lyrics.Language)
}
