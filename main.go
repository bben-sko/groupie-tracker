package main

import (
	"fmt"
	g "gt/func"
	"net/http"
)
func main() {
	http.HandleFunc("/", g.Home)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
