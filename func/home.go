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

func handleError(w http.ResponseWriter, status int, msg string, err error) {
	http.Error(w, msg, status)
	fmt.Println(err)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("s"))
	var results []d.SearchResult

	for i, artist := range artists {
		if len(results) > 16 {
			break
		}

		if i == 0 {
			for _, ar := range artists {
				artistName2 := strings.ToLower(ar.Name)
				if strings.HasPrefix(artistName2, query) {
					results = append(results, d.SearchResult{
						ID:   ar.ID,
						Name: ar.Name,
						Type: "artist/band",
					})
				}
			}
		}
		artistName := strings.ToLower(artist.Name)
		if !strings.HasPrefix(artistName, query) && strings.Contains(artistName, query) {
			results = append(results, d.SearchResult{
				ID:   artist.ID,
				Name: artist.Name,
				Type: "artist/band",
			})
		}
	}
	for i, artist := range artists {
		if len(results) > 16 {
			break
		}
		if i == 0 {
			for _, ar := range artists {
				for _, member := range ar.Members {
					artistName2 := strings.ToLower(member)
					if strings.HasPrefix(artistName2, query) {
						results = append(results, d.SearchResult{
							ID:   ar.ID,
							Name: member,
							Type: "member of " + ar.Name,
						})
					}
				}
			}
		}
		for _, member := range artist.Members {
			if !strings.HasPrefix(member, query) && strings.Contains(strings.ToLower(member), query) {
				results = append(results, d.SearchResult{
					ID:   artist.ID,
					Name: member,
					Type: "member of " + artist.Name,
				})
			}
		}
	}
	for _, artist := range artists {
		if len(results) > 16 {
			break
		}
		if strings.HasPrefix(strings.ToLower(artist.FirstAlbum), query) {
			results = append(results, d.SearchResult{
				ID:   artist.ID,
				Name: artist.FirstAlbum,
				Type: "FirstAlbum of " + artist.Name,
			})
		}

		C_Date := strconv.Itoa(artist.CreationDate)
		if strings.HasPrefix(strings.ToLower(C_Date), query) {
			results = append(results, d.SearchResult{
				ID:   artist.ID,
				Name: C_Date,
				Type: "Creation Date of " + artist.Name,
			})
		}
	}
	for _, artist := range artists {
		if len(results) > 16 {
			break
		}

		if err := fetchAndDecode(artist.Locations, &artis); err != nil {
			handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
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
	if len(results) > 16 {
		results = results[:16]
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handleError(w, http.StatusNotFound, "page not found 404", nil)
		return
	}

	if err := fetchAndDecode("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
		return
	}

	tmp, err := template.ParseFiles("template/home_page.html")
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
		return
	}

	if err := tmp.Execute(w, artists); err != nil {
		handleError(w, http.StatusInternalServerError, "Internal Server Error 500", err)
	}
}
