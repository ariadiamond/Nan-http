package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"fmt"
	"os/exec"
)

func ConstructIndex (w http.ResponseWriter, url string) {
	var file string
	if strings.HasSuffix(url, "/index") {
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
	top = bytes.ReplaceAll(top, []byte("<!--TITLE-->"), []byte("Aria's Notes"))
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
	folder := url[:strings.LastIndex(url, "/") + 1]

	// Build top
	top, _ := ioutil.ReadFile("Root/head.html")
	bottom, _ := ioutil.ReadFile("Root/footer.html")
	top = bytes.ReplaceAll(top, []byte("<!--TITLE-->"), []byte(files.title))
	top = bytes.ReplaceAll(top, []byte("<!--NAV-->"), []byte("<td><a href=\"index\">Index</a></td>"))
	
	scripts := ""
	for _, val := range(files.scripts) {
		scripts = scripts + "<script type=\"text/javascript\" src=\"" + val + "\"></script>\n"
	}
	top = bytes.ReplaceAll(top, []byte("<!--JS-->"), []byte(scripts))
	
	styles := ""
	for _, val := range(files.styles) {
		styles = styles + "<link rel=\"stylesheet\" href=\"" + val + "\">\n"
	}
	top = bytes.ReplaceAll(top, []byte("<!--STYLE-->"), []byte(styles))
	
	io.WriteString(w, string(top))
	
	for _, name := range(files.files) {
		var rest []byte
		var err error
		// Check type of file we are sending
		if strings.HasSuffix(name, ".md") { // convert from github flavored markdown
			cmd := exec.Command("pandoc", "-f", "gfm", "-t", "html", folder + name)
			rest, err = cmd.Output()
		} else {
			rest, err = ioutil.ReadFile(folder + name)
		}
		if err != nil {
			fmt.Println(err.Error())
			Whomst(w) // TODO: superfluous 404
			return
		}
		io.WriteString(w, string(rest))
	}

	io.WriteString(w, string(bottom))

}
