package main

import (
	"fmt"
	"net/http"
	"strconv"
)

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

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {

		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specif id", id)

}

func createSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		//w.WriteHeader(405)
		//w.Write([]byte("Method not allowed"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("this is for creating a snippet box"))
}