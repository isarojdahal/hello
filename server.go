package main

import (
	"fmt"
	"net/http"
)

func main() {

	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "About, World!")
	})

	http.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id:= r.URL.Query().Get("id")
		fmt.Fprintf(w, "User, World! %s", id)
		
	})



	http.ListenAndServe(":8080", nil)
}

