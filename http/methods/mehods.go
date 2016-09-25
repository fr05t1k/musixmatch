package methods

//Get an album from our database: name, release_date, release_type, cover art.
const AlbumGet string = "album.get"

//Get the list of all the tracks of an album.
const AlbumTracksGet string = "album.tracks.get"

//Get the discography of an artist.
const ArtistAlbumsGet string = "artist.albums.get"

//This api provides you the list of the top artists of a given country.
const ArtistChartGet string = "artist.chart.get"

//Get the artist data from our database.
const ArtistGet string = "artist.get"

//Search for artists in our database.
const ArtistSearch string = "artist.search"

//Match your track against our database.
const MatcherTrackGet string = "matcher.track.get"

//With this api you'll be able to get the base url for the tracking script
//you need to insert in your page to legalize your existent lyrics library.
const TrackingUrlGet string = "tracking.url.get"

//This api provides you the list of the top tracks of the supported
//countries.
const TrackChartGet string = "track.chart.get"

//Get a track info from our database: title, artist, instrumental flag and
//cover art.
const TrackGet string = "track.get"

//This API method provides you the opportunity to help us improving our
//catalogue. (v1.1)
const TrackLyricsFeedbackPost string = "track.lyrics.feedback.post"

//Retrieve the lyrics of a track.
const TrackLyricsGet string = "track.lyrics.get"

//Retrieve the lyric snippet of a track.
const TrackSnippetGet string = "track.snippet.get"

//Retrieve the subtitle of a track.
const TrackSubtitleGet string = "track.subtitle.get"

//Search for a track in our database.
const TrackSearch string = "track.search"

//Retrieve the subtitle of a track.
const TrackSubtitle string = "track.subtitle.get"
