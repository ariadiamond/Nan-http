package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Returns true if the url should NOT be accessed
// should be called isForbidden (TODO?)
func checkForbidden (url string, method string) (bool) {
	aclState, exist := ACL[url]
	if !exist { // If it doesn't exist, we are good
		return false
	}
	
	// It is restricted in some way, let's find out how
	if (aclState & NEVER) == NEVER {
		return true
	}
	if method == "GET" && ((aclState & SUREAD) == SUREAD && !SuRead) {
		return true
	}
	if method == "PUT" && (aclState & SUWRITE) == SUWRITE && !SuWrite {
		return true
	}
	if method == "PUT" && (aclState & READONLY) == READONLY {
		Error(method)
		return true
	}
	if method == http.MethodPut && !AllowPut {
		return true
	}
	
	// We made it past and didn't have any issues
	return false
}

func no (w http.ResponseWriter) {
	w.WriteHeader(403) // no (should I make a prettier screen for this?)
}

func Whomst (w http.ResponseWriter) {
	w.WriteHeader(404)
}

func Handle (w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if strings.Count(url, "/") == 1 {
		url = "./Root" + url
	} else {
		url = "." + url
	}

	// print access
	Info(r.Method, url)

	// check if the file is allowed
	if checkForbidden(url, r.Method) {
		no(w)
		return
	}
	
	switch(r.Method) {
	case http.MethodGet:
		Get(w, url)
	case http.MethodPut:
		Put(w, r, url)		
	default: // unsupported
		w.WriteHeader(405)
	}
}

func Get(w http.ResponseWriter, url string) {
	
	// we have an index
	if url[len(url) - 1] == '/' {
		url = url + "index"
	}
	// all pages we need to construct do not have a "." in the filename
	if strings.LastIndex(url, ".") < strings.LastIndex(url, "/") {
		ConstructPage(w, url)
		return
	}
	
	// Apple asks for special apple favicons, but I just am giving them the
	// regular png
	if strings.Contains(url, "icon") && strings.Contains(url, ".png"){
		file, _ := ioutil.ReadFile("Root/favicon.png")
		w.Header().Set("Content-Type", "image/png")
		io.WriteString(w, string(file))
		return
	}

	// Just a normal file that does not need to be constructed
	file, err := ioutil.ReadFile(url)
	if err != nil { // file not found
		Whomst(w)
		return
	}

	if strings.Contains(url, ".css") {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}
	w.WriteHeader(200)

	io.WriteString(w, string(file))
}

func Put(w http.ResponseWriter, r *http.Request, url string) {
	fileStat, err := os.Stat(url)
	if err != nil {
		// TODO, find a better way to check for a non-existent file
		_, err = os.OpenFile(url, os.O_WRONLY | os.O_CREATE | os.O_EXCL, 0644)
		if err != nil {
			w.WriteHeader(403)
			return
		}
		fileStat, _ = os.Stat(url)
	}
	
	// Assuming overwriting files usually makes them longer, will add override TODO
	if r.ContentLength < fileStat.Size() {
		w.WriteHeader(409) // Conflict
		return
	}
	
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	err = ioutil.WriteFile(url, body, 0644)
	if err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}
