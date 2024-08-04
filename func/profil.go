package gt

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type locations struct {
	Locations []string `json:"locations"`
	
}
type dates struct {
	Dates []string `json:"dates"`
}

/*
*
"id": 1,

	"datesLocations": {
	  "dunedin-new_zealand": [
	    "10-02-2020"
	  ],
*/
type Relation struct {
	//
	DatesLocations map[string]interface{} `json:"datesLocations"`
}

func Profil(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		fmt.Println(err)
		return
	}
	var local locations
	err = json.NewDecoder(response.Body).Decode(&local)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		fmt.Println(err)
		return
	}
	var date dates
	err = json.NewDecoder(res.Body).Decode(&date)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		fmt.Println(err)
		return
	}
	var relation Relation
	err = json.NewDecoder(resp.Body).Decode(&relation)
	if err != nil {
		fmt.Println(err)
		return
	}
	tmp, err := template.ParseFiles("template/profil_page.html")
	if err != nil {
		http.Error(w, "parse", http.StatusInternalServerError)
	}
	err = tmp.Execute(w, map[string]interface{}{
		"datalocal":    local,
		"datadate":     date,
		"datarelation": relation,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}
