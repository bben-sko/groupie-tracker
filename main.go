package main

import (
	"fmt"
	"net/http"

	g "gt/func"
)

func main() {
	http.Handle("/styles.css", http.FileServer(http.Dir("template")))
	http.HandleFunc("/", g.Home)
	http.HandleFunc("/search", g.SearchHandler)
	fmt.Println("http://localhost:8081/")

	http.HandleFunc("/profil", g.Profil)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
