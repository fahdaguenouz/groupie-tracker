package handlers

import (
	"groupie/models"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	template, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	customError := models.Error{
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	if err := template.Execute(w, customError); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
