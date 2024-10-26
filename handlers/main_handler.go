package handlers

import (
	"html/template"
	"net/http"
)


func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w,r,http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w,r,http.StatusMethodNotAllowed)
        return
	}
	template,err:=template.ParseFiles("templates/index.html")
	if err!=nil{
        http.Error(w,err.Error(),http.StatusInternalServerError)
        return
    }
	if err:=template.Execute(w,nil);err!=nil{
		ErrorHandler(w,r,http.StatusInternalServerError)
		return
	}

}
