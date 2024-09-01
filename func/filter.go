package handler

import (
	d "gt/data"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)



func Filter(w http.ResponseWriter, r *http.Request) {
	/*if r.Method != http.MethodPost {
		handleError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}
	if r.URL.Path != "/filter" {
		handleError(w, http.StatusNotFound, "page not found", nil)
		return
	}*/
	var results []d.Filter
	// Loop through all artists to find matching names and add them to results
	for i, artist := range artists {
		if Check_filter(r, artist.CreationDate, artist.Members, artist.FirstAlbum, i) {
			results = append(results, d.Filter{
				Image: artist.Image,
				ID:    artist.ID,
				Name:  artist.Name,
			})
		}
	}
	tmp, err := template.ParseFiles("template/Filter.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
		return
	}


	if results == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := tmp.Execute(w, results); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
	}
}

func Check_filter(r *http.Request, CreationDate int, Members []string, first_album string, i int) bool {
	creation_date_min, _ := strconv.Atoi(r.FormValue("creation_date_min"))
	creation_date_max, _ := strconv.Atoi(r.FormValue("creation_date_max"))
	first_album_date := r.FormValue("first_album_date")
	number_of_members, _ := strconv.Atoi(r.FormValue("number_of_members"))
	locUS := r.FormValue("locUS")
	locUK := r.FormValue("locUK")
	creation_date := false
	first_album_b := false
	members := false
	US := false
	UK := false

	if r.FormValue("creation_date_min") == "" && r.FormValue("creation_date_max") == "" {
		creation_date = true
	} else if CreationDate >= creation_date_min && CreationDate <= creation_date_max {
		creation_date = true
	}
	if first_album_date == "" {
		first_album_b = true
	} else if first_album == first_album_date {
		first_album_b = true
	}
	if r.FormValue("number_of_members") == "" {
		members = true
	} else {
		r := 0
		for range Members {
			r++
		}
		if r == number_of_members {
			members = true
		}else {
		members = false
		}
	}
	if locUS == "" && locUK == "" {
		US = true
		UK = true
	} else {
		for _, lo := range artis.Index[i].Locations {
			// Check if the query matches any location name
			if strings.HasSuffix(lo, locUS) {
				US = true
			}
			if strings.HasSuffix(lo, locUK) {
				UK = true
			}
		}
		if UK != true && US != true {
			US = false
			UK = false
		}
	}
	if  UK == true && US == true && members == true && first_album_b == true && creation_date == true {
		return true
	} else {
		return false
	}
}
