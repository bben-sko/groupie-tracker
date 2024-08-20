package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	d "gt/data"
)

var (
	artists []d.Artist
	artis   d.Locations
)

func handleError(w http.ResponseWriter, status int, msg string,err error) {
	http.Error(w, msg, status)
	fmt.Println(err)
}
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("s"))
	var results []d.SearchResult

	for _, artist := range artists {
		if len(results) > 5 {
			break
		}

		artistName := strings.ToLower(artist.Name)
		if strings.Contains(artistName, query) {
			results = append(results, d.SearchResult{
				ID:   artist.ID,
				Name: artist.Name,
				Type: "artist/band",
			})
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				results = append(results, d.SearchResult{
					ID:   artist.ID,
					Name: member,
					Type: "member of " + artist.Name,
				})
			}
		}

		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			results = append(results, d.SearchResult{
				ID:   artist.ID,
				Name: artist.FirstAlbum,
				Type: "FirstAlbum of " + artist.Name,
			})
		}

		C_Date := strconv.Itoa(artist.CreationDate)
		if strings.Contains(strings.ToLower(C_Date), query) {
			results = append(results, d.SearchResult{
				ID:   artist.ID,
				Name: C_Date,
				Type: "Creation Date of " + artist.Name,
			})
		}

		if err := fetchAndDecode(artist.Locations, &artis); err != nil {
			handleError(w, http.StatusInternalServerError, "Internal Server Error 500",err)
			return
		}

		for _, local := range artis.Locations {
			if strings.Contains(strings.ToLower(local), query) {
				results = append(results, d.SearchResult{
					ID:   artist.ID,
					Name: local,
					Type: "location of " + artist.Name,
				})
			}
		}
	}

	if len(results) > 5 {
		results = results[:5]
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500",err)
	}
}


func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, http.StatusNotFound, "page not found 404",nil)
		return
	}

	if err := fetchAndDecode("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500",err)
		return
	}

	tmp, err := template.ParseFiles("template/home_page.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500",err)
		return
	}

	if err := tmp.Execute(w, artists); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500",err)
	}
}
