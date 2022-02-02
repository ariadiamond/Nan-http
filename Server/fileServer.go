package main

import (
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

/* no handles when the client should not have access to this page, so deny them with a 403.
 * At the moment, this is unreachable code, but hopefully will be used again when I have a better
 * design for limiting access to certain files.
 */
func no (w http.ResponseWriter) {
    w.WriteHeader(403) // no (should I make a prettier screen for this?)
    ConstructPage(w, "./Root/accessDenied") // make sure this exists otherwise we will infinite loop
}

/* Whomst handles files that don't exist, so we deny with a 404 */
func Whomst (w http.ResponseWriter) {
    w.WriteHeader(404)
    ConstructPage(w, "./Root/fileNotFound") // make sure this exists otherwise we will infinite loop
}

/* Handle handles all requests, checking that certain requirements are handled, such as moving the
 * files in the root directory (as seen by the client), to Root/, as well as checking access 
 * permissions and finally passing the request based on the HTTP request type.
 */
func Handle (w http.ResponseWriter, r *http.Request) {
    url := r.URL.Path
    if strings.Count(url, "/") == 1 { // move requests to the root of the server to Root/
        url = "./Root" + url
    } else { // make the url a relative path, rather than absolute one
        url = "." + url
    }

    // print access
    Info(r.Method, url)

    // Based on the method, handle this differently
    switch(r.Method) {
    case http.MethodGet:
        Get(w, url)
    case http.MethodPut:
        Put(w, r, url)
    default: // unsupported
        w.WriteHeader(405)
    }
}

/* Get handles HTTP GET requests (surprise, I know). */
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

    // Just a normal file that does not need to be constructed
    file, err := ioutil.ReadFile(url)
    if err != nil { // file not found
        Whomst(w)
        return
    }

    if strings.HasSuffix(url, ".css") {
        w.Header().Set("Content-Type", "text/css; charset=utf-8")
    } else if strings.HasSuffix(url, ".js") {
        w.Header().Set("Content-Type", "script/javascript; charset=utf-8")
    }
    w.WriteHeader(200)

    w.Write(file)
}

/* Put handles put requests. At the current moment, it only supports writing new files. */
func Put (w http.ResponseWriter, r *http.Request, url string) {
    fileStat, err := os.Stat(url)
    if err != nil {
        fd, err := os.OpenFile(url, os.O_WRONLY | os.O_CREATE | os.O_EXCL, 0644)
        if err != nil {
            w.WriteHeader(403)
            return
        }
        fd.Close()
        fileStat, _ = os.Stat(url)
    }

    if r.ContentLength < fileStat.Size() {
        expectation, exist := r.Header["Expect"] // make Expect: 200 the override signal
        if !exist || (expectation[0] != "200") {
            w.WriteHeader(409) // conflict
            return
        }
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
