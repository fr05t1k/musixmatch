#MusixMatch
[![Build Status](https://travis-ci.org/fr05t1k/musixmatch.svg?branch=master)](https://travis-ci.org/fr05t1k/musixmatch)

A Go wrapper for the [MusixMatch API](https://developer.musixmatch.com/)

##How to use ?

#### Simple example
```go
package main

import (
    "github.com/fr05t1k/musixmatch"
    "fmt"
)
func main()  {
    mm := musixmatch.New("Your MusixMatch API Key") 
    lyrics, _ := mm.GetLyrics(113303287)
    
    fmt.Println(lyrics.Id)
    fmt.Println(lyrics.Body)
    fmt.Println(lyrics.Language)
}
```
#### Matching track
```
    track, _ := mm.GetMatchingTrack("Yesterday", "Beatles")
```
#### Search
```
    searchRequest := http.SearchRequest{Query:"Yesterday"}
    trackList, _ := mm.SearchTrack(searchRequest)
}
```

##Make custom request
You can also make custom request (which is not implemented)
```go

	params := url.Values{}
	params.Add(config.ApiKey, m.ApiKey)
	params.Add(config.QueryTrack, qTrack)
	params.Add(config.QueryArtist, qArtist)
	
	var trackResponse track.Response
	requestUrl := http.GetURLString(methods.MatcherTrackGet, params)
	err := http.RequestEntity(requestUrl, &trackResponse)
```
or
```go
	params := url.Values{}
	params.Add(config.TrackId, fmt.Sprintf("%d", trackId))
	resp, err := mm.Request(methods.TrackGet, params)

	var trackResponse track.Response

	json.Unmarshal(resp, &trackResponse)

	
```