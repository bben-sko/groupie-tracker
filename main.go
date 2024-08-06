package main

import (
	"fmt"
	"net/http"

	g "gt/func"
)

func main() {
	http.Handle("/style.css", http.FileServer(http.Dir("template")))
	http.HandleFunc("/", g.Home)
	http.HandleFunc("/profil", g.Profil)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("http://localhost:8081/")
}
