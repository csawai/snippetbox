package main

import (
	//"fmt"
	"log"
	"net/http"
	//"strconv"
)

func main() {

	mux := http.NewServeMux() // using servemux instead of http handlers as we get more flexibilty by assigning our own variable
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	//every error in go is a variable. In our case, error is running port :4000 to start server
	//it prints message else, you can
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
