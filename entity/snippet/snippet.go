package snippet

import "github.com/fr05t1k/musixmatch/entity"

type Response struct {
	Message Message `json:"message"`
}

type Message struct {
	Body   Body          `json:"body"`
	Header entity.Header `json:"header"`
}

type Body struct {
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	Language          string `json:"snippet_language"`
	Restricted        int8   `json:"restricted"`
	Instrumental      int32  `json:"instrumental"`
	Body              string `json:"snippet_body"`
	ScriptTrackingURL string `json:"script_tracking_url"`
	PixelTrackingURL  string `json:"pixel_tracking_url"`
	HtmlTrackingURL   string `json:"html_tracking_url"`
	UpdatedTime       string `json:"updated_time"`
}
