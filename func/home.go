package gt

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Assuming you have a struct like this for your data
// type Artist struct {
// 	ID           int
// 	Name         string
// 	CreationDate string
// 	FirstAlbum   string
// 	Image        string
// }

var artists = []Artist{
	// Your list of artists
}

type SearchResult struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Type string `json:"type"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    query := strings.ToLower(r.URL.Query().Get("s"))
    var results []SearchResult

    for _, artist := range artists {
        // Check if the query matches the artist/band name
        if strings.Contains(strings.ToLower(artist.Name), query) {
            results = append(results, SearchResult{
                ID:   artist.ID,
                Name: artist.Name,
                Type: "artist/band",
            })
        }
        
        // Check if the query matches any member's name
        for _, member := range artist.Members {
            if strings.Contains(strings.ToLower(member), query) {
                results = append(results, SearchResult{
                    ID:   artist.ID,
                    Name: member,
                    Type: "member",
                })
            }
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}


// func Home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "page not found 404", http.StatusNotFound)
// 		return
// 	}
// 	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
// 	if err != nil {
// 		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	// defer data_api.Body.Close()
// 	var art []Artist
// 	err = json.NewDecoder(resp.Body).Decode(&art)
// 	if err != nil {
// 		http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	// fmt.Println(art)
// 	tmp, err := template.ParseFiles("template/home_page.html")
// 	if err != nil {
// 		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
// 		return
// 	}
// 	err = tmp.Execute(w, map[string]interface{}{
// 		"data": art,
// 	})
// 	if err != nil {
// 		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
// 		return
// 	}
// }


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
    defer resp.Body.Close() // Don't forget to close the response body

    // Decode the fetched artist data
    err = json.NewDecoder(resp.Body).Decode(&artists)
    if err != nil {
        http.Error(w, "Status Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Render the home page template with the fetched data
    tmp, err := template.ParseFiles("template/home_page.html")
    if err != nil {
        http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
        return
    }
    err = tmp.Execute(w, map[string]interface{}{
        "data": artists,
    })
    if err != nil {
        http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
        return
    }
}
