package main

import (
	"fmt"
	"net/http"

	g "gt/func"
	handler "gt/func"
)

func servCss(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/assets/" {
		handler.ErrorPages(w, 404, "Page not found")
		return
	}
	fs := http.FileServer(http.Dir("./assets"))
	http.StripPrefix("/assets/", fs).ServeHTTP(w, r)
}

func main() {
	fmt.Println("http://localhost:8081/")
	http.HandleFunc("/", g.Home)
	http.HandleFunc("/search-query", g.SearchHandler)
	http.HandleFunc("/search", g.Search)
	http.HandleFunc("/profil", g.Profil)
	http.HandleFunc("/filter", g.Filter)

	http.HandleFunc("/assets/", servCss)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
