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
	creation_date_min, _  :=  strconv.Atoi(r.FormValue("creation_date_min"))
	creation_date_max, _  := strconv.Atoi(r.FormValue("creation_date_max"))
	first_album_date_min, _ := strconv.Atoi(r.FormValue("first_album_date_min"))
	first_album_date_max, _ :=  strconv.Atoi(r.FormValue("first_album_date_max"))
//	number_of_members := r.FormValue("number_of_members")
	locUS := r.FormValue("locUS")
	locUK := r.FormValue("locUK")

	var results []d.Filter

	// Loop through all artists to find matching names and add them to results
	for _, artist := range artists {
		fb := artist.FirstAlbum
		f_album, _ := strconv.Atoi(fb)
		if f_album >= first_album_date_min && f_album <= first_album_date_max {
			results = append(results, d.Filter{
				Image: artist.Image,
				ID:    artist.ID,
				Name:  artist.Name,
				//Type:  "FirstAlbum of " + artist.Name,
			})
		}

		// Check if the locUS matches the artist's creation date
		if artist.CreationDate >= creation_date_min && artist.CreationDate <= creation_date_max {
			results = append(results, d.Filter{
				Image: artist.Image,
				ID:    artist.ID,
				Name:   artist.Name,
				//Type:  "Creation Date of " + artist.Name,
			})
		}
	}

	// Loop through all artists again to find matching members and add them to results
/*	for i, artist := range artists {
		if len(results) > 16 {
			// Stop if we have reached the limit of 16 results
			break
		}
		if i == 0 {
			// On the first iteration, check if any member names start with the locUS
			for _, ar := range artists {
				for _, member := range ar.Members {
					artistName2 := strings.ToLower(member)
					if strings.HasPrefix(artistName2, locUS) {
						results = append(results, d.Filter{
							Image: ar.Image,
							ID:    ar.ID,
							Name:  member,
							Type:  "member of " + ar.Name,
						})
					}
				}
			}
		}

		// Check if any member names contain the locUS
		for _, member := range artist.Members {
			if !strings.HasPrefix(strings.ToLower(member), locUS) && strings.Contains(strings.ToLower(member), locUS) {
				results = append(results, d.Filter{
					Image: artist.Image,
					ID:    artist.ID,
					Name:  member,
					Type:  "member of " + artist.Name,
				})
			}
		}
	}
*/
	// Loop through all location indices to find matching locations and add them to results
	j := 0
	for _, loc := range artis.Index {
		name := artists[j].Name
		for _, lo := range loc.Locations {
			// Check if the locUS matches any location name
			if  strings.HasSuffix(strings.ToLower(lo), locUS) || strings.HasSuffix(strings.ToLower(lo), locUK) {
				if len(results) < 16 {
					results = append(results, d.Filter{
						Image: artists[j].Image,
						ID:    loc.ID,
						Name:  name,
						//Type:  "location " + name,
					})
				}
			}
		}
		j++
	}

	// Attempt to parse the search results template
	tmp, err := template.ParseFiles("template/Filter.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
		return
	}

	// Check if results are empty and handle accordingly
	if results == nil {
		// Attempt to parse the notfound template file
		tmp1, err := template.ParseFiles("template/notfound.html")
		if err != nil {
			// If template parsing fails, handle the error
			handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
			return
		}

		// Execute the notfound template
		err = tmp1.Execute(w, nil)
		if err != nil {
			handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
		}
		return
	}

	// Execute the search results template with the results
	if err := tmp.Execute(w, results); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
	}
}
