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
	if strings.Contains(url, "/2/") {
		return "Economics 2: Intro to Macroeconomics"
	} else if strings.Contains(url, "/115/") {
		return "CSE 115A: Intro to Software Engineering"
	} else if strings.Contains(url, "/180/") {
		return "CSE 180: Databases"
	} else if strings.Contains(url, "/290/") {
		return "CSE 290S: Resillience in Large Scale Systems"
	}
	//otherwise
	return "Winter 2021"
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
	lastIndex := strings.LastIndex(url, "/")
	var files []string
	if strings.Contains(url[lastIndex:], ".") {
		files = strings.Split(url[lastIndex + 1:], ".")
	} else {
		files = []string{url[lastIndex + 1:]}
	}

	// Build top
	top, _ := ioutil.ReadFile("Root/head.html")
	bottom, _ := ioutil.ReadFile("Root/footer.html")
	top = bytes.ReplaceAll(top, []byte("<!--TITLE-->"), []byte(getTitle(url) + " Notes"))
	top = bytes.ReplaceAll(top, []byte("<!--NAV-->"), []byte("<td><a href=\"./\">Index</a></td>"))
	io.WriteString(w, string(top))

	for _, name := range(files) {
		rest, err := ioutil.ReadFile(url[:lastIndex] + "/Notes/" + name + ".html")
		if err != nil {
			fmt.Println(err.Error())
			Whomst(w)
			return
		}
		io.WriteString(w, string(rest))
	}

	io.WriteString(w, string(bottom))

}
