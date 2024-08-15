package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	d "gt/data"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ID, _ := strconv.Atoi(id)
	if ID < 1 || ID > 52 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/profil" {
		http.Error(w, "page not found 404", http.StatusNotFound)
		return
	}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	var local d.Locations
	err = json.NewDecoder(response.Body).Decode(&local)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	var date d.Dates
	err = json.NewDecoder(res.Body).Decode(&date)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	res_ID_art, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	var artists_id d.Artist_id
	err = json.NewDecoder(res_ID_art.Body).Decode(&artists_id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	var relation d.Relation
	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
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
		"data_artist": artists_id,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
