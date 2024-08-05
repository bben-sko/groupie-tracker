package gt

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type locations struct {
	Locations []string `json:"locations"`
	
}
type dates struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	DatesLocations map[string]interface{} `json:"datesLocations"`
}

func Profil(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ID, _ := strconv.Atoi(id)
	if ID < 1 || ID > 52 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/profil" {
		http.Error(w, "page not found 404",http.StatusNotFound )
		return
	}
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	var local locations
	err = json.NewDecoder(response.Body).Decode(&local)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	var date dates
	err = json.NewDecoder(res.Body).Decode(&date)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	var relation Relation
	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	tmp, err := template.ParseFiles("template/profil_page.html")
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, map[string]interface{}{
		"datalocal":    local,
		"datadate":     date,
		"datarelation": relation,
	})
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
