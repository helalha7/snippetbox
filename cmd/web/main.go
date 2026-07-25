package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)                          // dsiplay the home page
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)     // display a specific snippet
	mux.HandleFunc("GET /snippet/create", snippetCreate)      // display a form for creating snippet
	mux.HandleFunc("POST /snippet/create", snippetCreatePost) // save a new snippet

	//stating the server
	fmt.Println("Starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
