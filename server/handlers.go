package server

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, r, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errors(w, r, http.StatusMethodNotAllowed)
		return
	}

	artists := grabjson.GetJsonData()
	files := []string{"ui/templates/index.html"}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		Errors(w, r, http.StatusInternalServerError)
		return
	}
	if err != nil {
		Errors(w, r, http.StatusInternalServerError)
		return
	}
}

func details(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/detail" {
		Errors(w, r, http.StatusNotFound)
		return
	}
	// checking query
	query := r.URL.Query()
	md, value := query["id"]
	if !value {
		Errors(w, r, http.StatusNotFound)
		log.Println("There is only [id] query value exists")
		return
	}
	var idValues []int
	for _, l := range md {
		j, err := strconv.Atoi(l)
		if err != nil {
			Errors(w, r, http.StatusBadRequest)
			log.Println("Probably inappropriate URL query")
			return
		}
		idValues = append(idValues, j)
	}
	for _, l := range idValues {
		if l < 1 || l > 52 {
			Errors(w, r, http.StatusBadRequest)
			log.Println("There are only 52 artists")
			return
		}
	}
	if r.Method != http.MethodGet {
		Errors(w, r, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		Errors(w, r, http.StatusBadRequest)
		return
	}
	artist := grabjson.GetJsonData()
	detailArtist := grabjson.GetMapData(id, &artist[id-1])
	files := []string{"ui/templates/detail.html"}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		Errors(w, r, http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "detail.html", detailArtist)
	if err != nil {
		Errors(w, r, http.StatusInternalServerError)
		return
	}
}
