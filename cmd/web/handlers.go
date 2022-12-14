package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{

		"ui/html/home.page.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
	}

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
