package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	d "gt/data"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}
	if r.URL.Path != "/filter" {
		handleError(w, http.StatusNotFound, "page not found", nil)
		return
	}
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
	first_album_date_min, _ := strconv.Atoi(r.FormValue("first_album_min"))
	first_album_date_max, _ := strconv.Atoi(r.FormValue("first_album_max"))
	var str []int
	for i := 1; i < 9; i++ {
		m, err := strconv.Atoi(r.FormValue("num_members_" + strconv.Itoa(i)))
		if err != nil {
			continue
		}
		str = append(str, m)
	}
	locUK := r.FormValue("city")

	if r.FormValue("creation_date_min") != "" || r.FormValue("creation_date_max") != "" {
		if r.FormValue("creation_date_min") == "" {
			creation_date_min = 1970
		}
		if r.FormValue("creation_date_max") == "" {
			creation_date_max = 2024
		}

		if !(CreationDate >= creation_date_min && CreationDate <= creation_date_max) {
			return false
		}
	}

	if r.FormValue("first_album_min") != "" || r.FormValue("first_album_max") != "" {
		if r.FormValue("first_album_min") == "" {
			first_album_date_min = 1978
		}
		if r.FormValue("first_album_max") == "" {
			first_album_date_max = 2024
		}

		first_album_date, err := strconv.Atoi(first_album[6:])
		if err != nil {
			return false
		}
		if !(first_album_date >= first_album_date_min && first_album_date <= first_album_date_max) {
			return false
		}
	}

	numMembers, err := strconv.Atoi(r.FormValue("members"))
	if err != nil {
		return false
	}
	if len(Members) != numMembers {
		return false
	}

	if str != nil {
		r := 0
		for range Members {
			r++
		}
		if !Is_here(str, r) {
			return false
		}
	}
	if locUK != "" {
		k := 0
		for _, lo := range artis.Index[i].Locations {
			if strings.HasSuffix(lo, locUK) {
				k++
			}
		}
		if k == 0 {
			return false
		}
	}
	return true
}

func Is_here(str []int, r int) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == r {
			return true
		}
	}
	return false
}
