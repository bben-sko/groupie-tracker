package main

import (
	"fmt"
	"net/http"

	g "gt/func"
)

func main() {
	fmt.Println("http://localhost:8080/")
	http.HandleFunc("/", g.Home)
	http.HandleFunc("/searc", g.SearchHandler)
	http.HandleFunc("/search", g.Search)


	http.HandleFunc("/profil", g.Profil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
