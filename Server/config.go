package main

import (
    "io/ioutil"
    "regexp"
    "strings"
)

var Config map[string](map[string]ConfVal)

type ConfVal struct {
    title   string
    files   []string
    scripts []string
    styles  []string
}

// defConf exists so that we can return a ConfVal, rather than pointer.
var defConf = ConfVal{title: "File Not found", files: nil, scripts: nil, styles: nil}

/* Readconfig gets config information for a specific URL to help the page to be displayed. It
 * essentially makes the Config variable act like a cache, where the file is retrieved if it does
 * not exist in memory (a cache miss), or returns the data (a cache hit). This also puts config
 * information in the case of non-existent config files to prevent future failures.
 */
func ReadConfig (url string) (ConfVal, bool) {
    lastIndex := strings.LastIndex(url, "/")
    folder    := url[:lastIndex + 1]
    file      := url[lastIndex + 1:]
    folderConfig, exist := Config[folder]
    if !exist { // If we don't already have it, try to get it
        if !parseConfig(folder) {
            return defConf, false
        }
        folderConfig, _ = Config[folder]
    }

    files, exist := folderConfig[file]
    if !exist {
        return defConf, false
    }
    return files, true
}

/* parseConfig takes a folder, gets the config file and parses it. This function removes comments
 * and separates entries, but passes the majority of the parsing work to parseLine.
 */
func parseConfig (folder string) bool {
    contents, err := ioutil.ReadFile(folder + ".httpconfig")
    if err != nil { // so we don't parse it again
        Config[folder] = make(map[string]ConfVal)
        return false
    }

    list       := strings.Split(string(contents), "*")
    fileConfig := make(map[string]ConfVal)
    // check for comments (#)
    for _, line := range(list) {
        cmtIdx := regexp.MustCompile(`(?m:(#)([[:blank:]]|[[:graph:]])*$)`)
        for cmtSE := cmtIdx.FindIndex([]byte(line));
            cmtSE != nil;
            cmtSE = cmtIdx.FindIndex([]byte(line)) {
            line = line[:cmtSE[0]] + line[cmtSE[1]:]
        }
        if len(line) == 0 {
            continue
        }
        url, files := parseLine(line)
        if len(url) == 0 {
            Warn("Unable to parse:\n" + line)
        }
        fileConfig[url] = files
    }
    Config[folder] = fileConfig
    return true
}

/* parseLine does the hard work of parsing an individual entry. */
func parseLine(line string) (string, ConfVal) {
    endTitle := strings.Index(line, "@")
    var confVal ConfVal
    if endTitle == -1 { // no title (which is okay)
        confVal.title = "Aria's notes" // TODO: make file specific default
    } else { // found title
        confVal.title = strings.TrimSpace(line[:endTitle])
        line = line[endTitle + 1:]
    }
    urlRE := regexp.MustCompile(`(\s*)(=>)`)
    urlSE := urlRE.FindIndex([]byte(line))
    if urlSE == nil {
        //error
        return "", confVal
    }
    // create + parse regex
    scriptRE := regexp.MustCompile(`(?i:scripts)(\s*)(=>)`)
    scriptSE := scriptRE.FindIndex([]byte(line))

    styleRE := regexp.MustCompile(`(?i:styles)(\s*)(=>)`)
    styleSE := styleRE.FindIndex([]byte(line))

    // parse actual pieces so we can create the things
    url := strings.TrimSpace(line[:urlSE[0]])
    filesEnd := len(line)
    if scriptSE != nil && scriptSE[0] < filesEnd {
        filesEnd = scriptSE[0]
    }
    if styleSE != nil && styleSE[0] < filesEnd {
        filesEnd = styleSE[0]
    }
    files := strings.Split(line[urlSE[1]:filesEnd], ",")
    confVal.files = make([]string, len(files))
    // we need error checking to have non-zero lengths?
    for idx, val := range(files) {
        confVal.files[idx] = strings.TrimSpace(val)
    }

    if scriptSE != nil {
        end := len(line)
        if urlSE[1] > scriptSE[1] {
            end = urlSE[0]
        }
        if styleSE != nil && styleSE[1] > scriptSE[1] && styleSE[0] < end {
            end = styleSE[0]
        }
        scripts := strings.Split(line[scriptSE[1] + 1:end], ",")
        confVal.scripts = make([]string, len(scripts))
        for idx, val := range(scripts) {
            confVal.scripts[idx] = strings.TrimSpace(val)
        }
    }

    if styleSE != nil {
        end := len(line)
        if urlSE[1] > styleSE[1] {
            end = urlSE[0]
        }
        if scriptSE != nil && scriptSE[1] > styleSE[1] && scriptSE[0] < end {
            end = scriptSE[0]
        }
        styles := strings.Split(line[styleSE[1] + 1:end], ",")
        confVal.styles = make([]string, len(styles))
        for idx, val := range(styles) {
            confVal.styles[idx] = strings.TrimSpace(val)
        }
    }

    return url, confVal
}
