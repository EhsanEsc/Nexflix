package controller

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Register(tps *template.Template) {
	hc := new(homeController)
	hc.template = tps.Lookup("home.html")
	http.HandleFunc("/home", hc.get)

	// Use serveResource when using another mux
	// http.HandleFunc("/css/", serveResource)
	http.Handle("/css/", http.FileServer(http.Dir("resources")))
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "./resources" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content Type", contentType)

		bufferedReader := bufio.NewReader(f)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}
