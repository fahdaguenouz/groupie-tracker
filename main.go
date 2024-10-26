package main

import (
	"fmt"
	"groupie/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("/",handlers.MainHandler)
	http.HandleFunc("/artist/{id}",handlers.MainHandler)
	http.HandleFunc("/assets/", handlers.AssetsHandler)

	fmt.Println("Server is running at http://localhost:3000")
	err:=http.ListenAndServe(":3000",nil)
	if err!=nil{
        fmt.Println("Error starting server: ",err)
    }

}