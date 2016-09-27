package list

import (
	"github.com/fr05t1k/musixmatch/entity"
	"github.com/fr05t1k/musixmatch/entity/track"
)

type Response struct {
	Message Message `json:"message"`
}

func (r *Response) GetMessage() Message {
	return r.Message
}

type Message struct {
	Body   Body          `json:"body"`
	Header entity.Header `json:"header"`
}

type Body struct {
	Tracks []TrackList `json:"track_list"`
}

type TrackList struct {
	Track track.Track `json:"track"`
}
