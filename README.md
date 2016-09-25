#MusixMatch
[![Build Status](https://travis-ci.org/fr05t1k/musixmatch.svg?branch=master)](https://travis-ci.org/fr05t1k/musixmatch)

A Go wrapper for the [MusixMatch API](https://developer.musixmatch.com/)

##How to use ?

* Declaring the MusixMatch Instance

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
