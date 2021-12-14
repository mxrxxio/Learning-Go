package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
    // "reflect"
)

func main() {
    fetch("https://rss.applemarketingtools.com/api/v2/us/music/most-played/10/albums.json")
}

type Album struct {
	Feed struct {
		Title  string `json:"title"`
		ID     string `json:"id"`
		Author struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"author"`
		Links []struct {
			Self string `json:"self"`
		} `json:"links"`
		Copyright string `json:"copyright"`
		Country   string `json:"country"`
		Icon      string `json:"icon"`
		Updated   string `json:"updated"`
		Results   []struct {
			ArtistName            string `json:"artistName"`
			ID                    string `json:"id"`
			Name                  string `json:"name"`
			ReleaseDate           string `json:"releaseDate"`
			Kind                  string `json:"kind"`
			ArtistID              string `json:"artistId"`
			ArtistURL             string `json:"artistUrl"`
			ContentAdvisoryRating string `json:"contentAdvisoryRating,omitempty"`
			ArtworkURL100         string `json:"artworkUrl100"`
			Genres                []struct {
				GenreID string `json:"genreId"`
				Name    string `json:"name"`
				URL     string `json:"url"`
			} `json:"genres"`
			URL string `json:"url"`
		} `json:"results"`
	} `json:"feed"`
}

func fetch(url string) {
    res, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    data, _ := ioutil.ReadAll(res.Body)
    var title Album
    if err = json.Unmarshal([]byte(data), &title); err != nil {
        log.Fatal(err)
    }
    results := title.Feed.Results // struct
    for i := 0; i < len(results); i++ {
        fmt.Println(title.Feed.Results[i].ArtistName, " - ", title.Feed.Results[i].Name)
    }
}
