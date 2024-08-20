package main

import (
	"fmt"
	"net/http"

	g "gt/func"
)

func main() {
	fmt.Println("http://localhost:8080/")
	http.Handle("/styles.css", http.FileServer(http.Dir("template")))
	http.HandleFunc("/", g.Home)
	http.HandleFunc("/search", g.SearchHandler)

	http.HandleFunc("/profil", g.Profil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
