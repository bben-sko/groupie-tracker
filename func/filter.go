package handler

import (
	"fmt"
	"net/http"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	creation_date_min := r.FormValue("creation_date_min")
	creation_date_max := r.FormValue("creation_date_max")
	first_album_date_min := r.FormValue("first_album_date_min")
	first_album_date_max := r.FormValue("first_album_date_max")
	number_of_members := r.FormValue("number_of_members")
	locUS := r.FormValue("locUS")
	locUK := r.FormValue("locUK")
	fmt.Println(locUS,locUK)
	fmt.Println(creation_date_min,creation_date_max,first_album_date_min,first_album_date_max,number_of_members)
}