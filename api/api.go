package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    Relations
}

type Relations struct {
	ID             int `json:"id"`
	DatesLocations map[string][]string
}

const (
	url = "https://groupietrackers.herokuapp.com/api"
)

func GetArtistData() []Artist {
	r, err := http.Get(url + "/" + "artists")
	if err != nil {
		log.Println("can not get artists url")
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error in reading json data", err)
	}
	var artist []Artist
	err = json.Unmarshal(data, &artist)
	return artist
}

func GetDetailedData(i int, u *Artist) Artist {
	t, err := http.Get(url + "/" + "relation" + "/" + strconv.Itoa(i))
	if err != nil {
		log.Println("can not get Relations url")
	}
	defer t.Body.Close()
	dataTemp, err := io.ReadAll(t.Body)
	if err != nil {
		log.Println("Error in reading detail of the artist", err)
	}
	var relations Relations
	err = json.Unmarshal(dataTemp, &relations)
	if err != nil {
		log.Println("Error in unmarshalling json data in detail of the artist", err)
	}
	u.Relations = relations
	return *u
}
