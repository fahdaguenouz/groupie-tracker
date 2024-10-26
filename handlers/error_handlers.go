package handlers

import (
	"html/template"
	"net/http"
)

type Error struct {
	StatusCode int
	Error      string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request,statusCode int) {
	template,err:=template.ParseFiles("templates/error.html")
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
        return
	}
	customError:=Error{
		StatusCode: http.StatusNotFound,
        Error:      http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	if err:=template.Execute(w,customError);err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
        return
	}

}