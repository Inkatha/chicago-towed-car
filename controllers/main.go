package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

// Register handles routing for the application
func Register(templates *template.Template) {
	router := mux.NewRouter()

	homeController := new(homeController)
	homeController.template = templates.Lookup("home.html")
	router.HandleFunc("/", homeController.get)

	http.Handle("/", router)

	http.HandleFunc("/images/", serveResource)
	http.HandleFunc("/css/", serveResource)
}

// serveResource send css and image files to the server
func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err != nil {
		w.WriteHeader(404)
	}

	defer f.Close()

	w.Header().Add("Content-Type", contentType)

	br := bufio.NewReader(f)
	br.WriteTo(w)
}
