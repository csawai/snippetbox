package main

import (
	//"fmt"
	"log"
	"net/http"
)

//home function is the route to the home page. This will get triggered when main function is called.  We want to show some
//message of what happens when we visit this page. This page returns a string
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello snippetboxer"))

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is a show box"))

}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("this is for creating a snippet box"))
}

func main() {

	mux := http.NewServeMux() // using servemux instead of http handlers as we get more flexibilty by assigning our own variable
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	//every error in go is a variable. In our case, if error is port :4000 server start,
	//it prints message else, you can
	log.Fatal(err)
}
