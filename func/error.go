package handler

import (
	"html/template"
	"net/http"
)

type ErrorType struct {
	ErrorCode int
	Message   string
}

func ErrorPages(w http.ResponseWriter, code int, message string) {
	t, err := template.ParseFiles("template/error.html")
	if err != nil {
		t.Execute(w, ErrorType{
			ErrorCode: 500,
			Message:   "internal server error",
		})
		return
	}
	t.Execute(w, ErrorType{ErrorCode: code, Message: message})
}
