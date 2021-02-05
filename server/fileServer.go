package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"strings"
)

func checkForbidden (url string) (bool) {
	_, exist := Forbidden[url]
	if !exist {
		return false
	}
	return true
}

func no (w http.ResponseWriter) {
	w.WriteHeader(403) // no (should I make a prettier screen for this?)
}

func Whomst (w http.ResponseWriter) {
	w.WriteHeader(404)
}

func Handle (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}

	// Change this to more prettyness
	Warn("[GET] " + r.URL.Path)

	// check if the file is allowed
	if checkForbidden(r.URL.Path) {
		no(w)
		return
	}

	// all notes that need to be constructed have a _
	if strings.Contains(r.URL.Path, "_") {
		ConstructNotes(w, "." + r.URL.Path)
		return
	} else if strings.Contains(r.URL.Path, "index") || r.URL.Path[len(r.URL.Path) - 1] == '/' { // index is in index
		ConstructIndex(w, "." + r.URL.Path)
		return
	}

	// Apple asks for special apple favicons, but I just am giving them the regular png
	if strings.Contains(r.URL.Path, "icon") && strings.Contains(r.URL.Path, ".png"){
		file, _ := ioutil.ReadFile("favicon.png")
		w.Header().Set("Content-Type", "image/png")
		io.WriteString(w, string(file))
		return
	}

	// Just a normal file that does not need to be constructed
	file, err := ioutil.ReadFile("." + r.URL.Path)
	if err != nil { // file not found
		Whomst(w)
		return
	}

	if strings.Contains(r.URL.Path, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}
	w.WriteHeader(200)

	io.WriteString(w, string(file))
}
