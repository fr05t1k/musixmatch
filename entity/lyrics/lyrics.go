package lyrics

import "github.com/fr05t1k/musixmatch/entity"

type Response struct {
	Message Message `json:"message"`
}
type Message struct {
	Header header.Header `json:"header"`
	Body   LyricsBody    `json:"body"`
}
type LyricsBody struct {
	Lyrics Lyrics `json:"lyrics"`
}

type Lyrics struct {
	Id                uint32 `json:"lyrics_id"`
	Body              string `json:"lyrics_body"`
	Copyright         string `json:"lyrics_copyright"`
	Language          string `json:"lyrics_language"`
	PixelTrackingUrl  string `json:"pixel_tracking_url"`
	ScriptTrackingURL string `json:"script_tracking_url"`
}
