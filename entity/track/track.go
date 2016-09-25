package track

type Track struct {
	Id uint32 `json:"track_id"`
	// A MusicBrainz track identifier.
	MbId string `json:"track_mbid"`
	// A Spotify track identifier
	SpId string `json:"track_spotify_id"`
	// A Soundcloud track identifier
	ScId                 string `json:"track_soundcloud_id"`
	Rating               int32  `json:"track_rating"`
	Length               int32  `json:"track_length"`
	CommontrackId        uint32 `json:"commontrack_id"`
	Explicit             uint32 `json:"explicit"`
	HasLyrics            bool   `json:"has_lyrics"`
	HasSubtitles         bool   `json:"has_subtitles"`
	NumFavorite          uint32 `json:"num_favourite"`
	AlbumCoverArt100x100 string `json:"album_coverart_100x100"`
	AlbumCoverArt350x350 string `json:"album_coverart_350x350"`
	AlbumCoverArt500x500 string `json:"album_coverart_500x500"`
	AlbumCoverArt800x800 string `json:"album_coverart_800x800"`
	EditUrl              string `json:"track_edit_url"`
	UpdatedTime          string `json:"updated_time"`
	AlbumId              string `json:"album_id"`
	ArtistId             string `json:"album_name"`
	ArtistMbId           string `json:"artist_mbid"`
	ArtistName           string `json:"artist_name"`
	ShareUrl             string `json:"track_share_url"`
	Instrumental         int32  `json:"instrumental"`
	LyricsId             uint32 `json:"lyrics_id"`
	LyricsLength         uint32 `json:"lyrics_length"`
	SubtitleId           uint32 `json:"subtitle_id"`
	Name                 string `json:"track_name"`
}
