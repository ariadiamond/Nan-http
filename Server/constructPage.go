package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
)

func getTitle (url string) (string) {
	// TODO
	return "Spring 2021"
}

func ConstructIndex (w http.ResponseWriter, url string) {
	var file string
	if strings.Contains(url, "index.html") {
		file = url
	} else if strings.Contains(url, "index") {
		file = url + ".html"
	} else if url[len(url) - 1] == '/' {
		file = url + "index.html"
	} else { // There shouldn't be any cases
		Error("Construct Index has invalid string" + url)
	}
	rest, err := ioutil.ReadFile(file)
	if err != nil {
		Whomst(w)
		return
	}

	top, _ := ioutil.ReadFile("Root/head.html")
	bottom, _ := ioutil.ReadFile("Root/footer.html")
	top = bytes.ReplaceAll(top, []byte("<!--TITLE-->"), []byte(getTitle(url)))
	io.WriteString(w, string(top))
	io.WriteString(w, string(rest))
	io.WriteString(w, string(bottom))
}

func ConstructNotes (w http.ResponseWriter, url string) {
	files, exist := ReadConfig(url)
	if !exist { // we don't have a config for this url
		Whomst(w)
		return
	}

	// Build top
	top, _ := ioutil.ReadFile("Root/head.html")
	bottom, _ := ioutil.ReadFile("Root/footer.html")
	top = bytes.ReplaceAll(top, []byte("<!--TITLE-->"), []byte(getTitle(url) + " Notes"))
	top = bytes.ReplaceAll(top, []byte("<!--NAV-->"), []byte("<td><a href=\"../\">Index</a></td>"))
	io.WriteString(w, string(top))

	folder := url[:strings.LastIndex(url, "/") + 1]
	for _, name := range(files) {
		rest, err := ioutil.ReadFile(folder + name)
		if err != nil {
			fmt.Println(err.Error())
			Whomst(w) // TODO: superfluous 404
			return
		}
		io.WriteString(w, string(rest))
	}

	io.WriteString(w, string(bottom))

}
