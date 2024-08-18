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

var artists = []d.Artist{
	// Your list of artists
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("s"))
	var results []d.SearchResult

	for _, artist := range artists {

		if strings.Contains(strings.ToLower(artist.Name), query) {
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
					Type: "member",
				})
			}
		}
		
			if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
				results = append(results, d.SearchResult{
					ID:   artist.ID,
					Name: artist.FirstAlbum,
					Type: "FirstAlbum",
				})
			}
			C_Date := strconv.Itoa(artist.CreationDate) 
			
		if strings.Contains(strings.ToLower(C_Date), query) {
				results = append(results, d.SearchResult{
					ID:   artist.ID,
					Name: C_Date,
					Type: "Creation Date",
				})
			}
		
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found 404", http.StatusNotFound)
		return
	}
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	tmp, err := template.ParseFiles("template/home_page.html")
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = tmp.Execute(w, artists)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
