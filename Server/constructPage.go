package main

import (
    "bytes"
    "io/ioutil"
    "net/http"
    "os"
    "os/exec"
    "strings"
)

/* ConstructPage takes a URL and active connection, then builds the page (provided it has been
 * provided in a valid configuration file). It adds style and javascript links as specified, and
 * adds a header and end of the HTML page so it does not need to be included in each file. This is
 * useful when the same page is included in multiple pages (such as all notes and a specific topic,
 * as I use it), and there are not worries about including multiple HTML header sections. 
 *
 * This also converts any Markdown files to HTML included in the config, provided Pandoc is
 * installed. If Pandoc is not installed, any pages with Markdown files will fail to render,
 * returning a 404 response code.
 */
func ConstructPage (w http.ResponseWriter, url string) {
    files, exist := ReadConfig(url)
    if !exist { // we don't have a config for this url
        Whomst(w)
        return
    }
    
    // check if we have a cache before trying to rebuild it
    cache, err := ioutil.ReadFile(url + ".nancache")
    if err == nil {
        w.Write(cache)
        return
    }
    
    folder := url[:strings.LastIndex(url, "/") + 1]

    // Build top
    top, _ := ioutil.ReadFile("Root/head.html")
    bottom, _ := ioutil.ReadFile("Root/footer.html")
    top = bytes.ReplaceAll(top, []byte("<!--TITLE-->"), []byte(files.title))
    top = bytes.ReplaceAll(top,
                           []byte("<!--NAV-->"),
                           []byte("<td><a href=\"../index\">Index</a></td>"))

    // Include JavaScript and CSS files
    scripts := ""
    for _, val := range(files.scripts) {
        scripts = scripts + "<script type=\"application/javascript\" src=\"" + val + "\" defer></script>\n"
    }
    top = bytes.ReplaceAll(top, []byte("<!--JS-->"), []byte(scripts))

    styles := ""
    for _, val := range(files.styles) {
        styles = styles + "<link rel=\"stylesheet\" href=\"" + val + "\">\n"
    }
    top = bytes.ReplaceAll(top, []byte("<!--STYLE-->"), []byte(styles))

    // Build cache file so we can not superfluous 404
    // TODO, make caches exist until the server is closed, or reset removes it?
    // TODO would it be too much to ask to check if the file has been changed since creating the
    //      cached file
    tmp, _ := os.Create(url + ".nancache")

    tmp.Write(top)

    // Now iterate through the files in the config to read them and possibly convert
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

        if err != nil { // this means the file doesn't exist or the conversion failed
            Error(err.Error()) // Because this is a config file, we expect it to work, so rather
                               // than ignoring the failure (such as if someone tried to get a
                               // non-existent file, we print it.
            Whomst(w) // 404
            return
        }
        tmp.Write(rest)
    }

    // finish HTML file
    tmp.Write(bottom)

    // read the cache file and send it to the client
    body, _ := ioutil.ReadFile(tmp.Name())
    w.Write(body)

}
