package main

import (
	"fmt"
	"net/http"

	g "gt/func"
)

func main() {
	fmt.Println("http://localhost:8081/")
	http.HandleFunc("/", g.Home)
	http.HandleFunc("/search-query", g.SearchHandler)
	http.HandleFunc("/search", g.Search)
	http.HandleFunc("/profil", g.Profil)

	fs := http.FileServer(http.Dir("./template"))
    http.Handle("/template/", http.StripPrefix("/template/", fs))
	http.Handle("/style.css", http.FileServer(http.Dir("template")))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
