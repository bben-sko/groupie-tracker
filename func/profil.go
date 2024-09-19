package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	d "gt/data"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPages(w, 500, "internal server error")
		return
	}
	// Initialize a slice to hold split string parts
	var str []string

	// Get the 'id' query parameter from the URL
	id := r.URL.Query().Get("id")

	// Convert the 'id' from string to integer
	ID, _ := strconv.Atoi(id)

	// Split the 'id' by '/' to check its structure
	str = strings.Split(id, "/")

	// Check if the 'id' contains more than one part
	if len(str) > 1 {
		ErrorPages(w, 404,"not found")
		return
	}

	// Validate the 'ID' to ensure it is within the valid range
	if ID < 1 || ID > 52 {
		tmp1, err := template.ParseFiles("template/error.html")
		if err != nil {
			// If template parsing fails, handle the error
			ErrorPages(w, 500, "internal server error")
			return
		}

		// Execute the notfound template
		err = tmp1.Execute(w, nil)
		if err != nil {
			ErrorPages(w, 404, "not found")
		}
		return

	}

	// Base URL for the API requests
	baseURL := "https://groupietrackers.herokuapp.com/api"

	// Initialize variables to store API responses
	var local d.Locations
	var date d.Dates
	var artists_id d.Artist
	var relation d.Relation

	// Map of API endpoints to their corresponding response variables
	endpoints := map[string]interface{}{
		"/locations/": &local,
		"/dates/":     &date,
		"/artists/":   &artists_id,
		"/relation/":  &relation,
	}

	// Fetch and decode data from each API endpoint
	for endpoint, target := range endpoints {
		err := fetchAndDecode(baseURL+endpoint+id, target)
		if err != nil {

			ErrorPages(w, 500, "internal server error")
			return
		}
	}

	// Parse the HTML template for the profile page
	tmp, err := template.ParseFiles("template/profil_page.html")
	if err != nil {

		ErrorPages(w, 500, "internal server error")
		return
	}

	// Execute the template with data to render the profile page
	err = tmp.Execute(w, map[string]interface{}{
		"datalocal":    local,
		"datadate":     date,
		"datarelation": relation,
		"data_artist":  artists_id,
	})
	if err != nil {
		// fmt.Println(err)
		ErrorPages(w, 500, "internal server error")
		return
	}
}
