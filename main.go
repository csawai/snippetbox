package main

import (
	//"fmt"
	"log"
	"net/http"
)

//home function is the route to the home page. This will get triggered when main function is called.  We want to show some
//message of what happens when we visit this page. This page returns a string
func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("hello snippetboxer"))

	//check the URL package for this command

}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is a show box"))

}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	}

	w.Write([]byte("this is for creating a snippet box"))
}

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
