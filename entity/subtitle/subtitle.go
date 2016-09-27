package subtitle

import (
	"github.com/fr05t1k/musixmatch/entity"
)

type Response struct {
	Message Message `json:"message"`
}

type Message struct {
	Body   Body          `json:"body"`
	Header entity.Header `json:"header"`
}

type Body struct {
	Subtitle Subtitle `json:"subtitle"`
}

type Subtitle struct {
	Id                uint32 `json:"subtitle_id"`
	Restricted        int8   `json:"restricted"`
	Body              string `json:"subtitle_body"`
	Language          string `json:"subtitle_language"`
	PixelTrackingURL  string `json:"pixel_tracking_url"`
	ScriptTrackingURL string `json:"script_tracking_url"`
	HtmlTrackingURL   string `json:"html_tracking_url"`
}
