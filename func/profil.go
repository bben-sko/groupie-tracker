package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	d "gt/data"
)




func Profil(w http.ResponseWriter, r *http.Request) {
	var str []string
	id := r.URL.Query().Get("id")
	ID, _ := strconv.Atoi(id)
	str = strings.Split(id, "/")

	if len(str) > 1 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if ID < 1 || ID > 52 {
		http.Error(w, "bad request ghghgh", http.StatusBadRequest)
		return
	}

	baseURL := "https://groupietrackers.herokuapp.com/api"
	var local d.Locations
	var date d.Dates
	var artists_id d.Artist
	var relation d.Relation

	endpoints := map[string]interface{}{
		"/locations/": &local,
		"/dates/":     &date,
		"/artists/":   &artists_id,
		"/relation/":  &relation,
	}

	for endpoint, target := range endpoints {
		err := fetchAndDecode(baseURL+endpoint+id, target)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal Server Error 500 ", http.StatusInternalServerError)
			return
		}
	}

	tmp, err := template.ParseFiles("template/profil_page.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, map[string]interface{}{
		"datalocal":    local,
		"datadate":     date,
		"datarelation": relation,
		"data_artist":  artists_id,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
