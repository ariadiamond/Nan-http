package main

import (
	"io/ioutil"
	"strings"
)

var Config map[string](map[string]ConfVal)
type ConfVal struct {
	title string
	files []string
}
var defConf = ConfVal{title: "File Not found", files: nil}

func ReadConfig(url string) (ConfVal, bool) {
	lastIndex := strings.LastIndex(url, "/")
	folder    := url[:lastIndex + 1]
	file      := url[lastIndex + 1:]
	folderConfig, exist := Config[folder]
	if !exist { // If we don't already have it, try to get it
		if !ParseConfig(folder) {
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

func ParseConfig(folder string) (bool) {
	contents, err := ioutil.ReadFile(folder + ".httpconfig")
	if err != nil { // so we don't parse it again
		Config[folder] = make(map[string]ConfVal)
		return false
	}

	list       := strings.Split(string(contents), "\n")
	fileConfig := make(map[string]ConfVal)
	// check for comments (#)
	for _, val := range(list) {
		if len(val) == 0 {
			continue
		}
		cmtIdx := strings.Index(val, "#")
		var line string
		if cmtIdx > 0 {
			line = val[:cmtIdx]
		} else if cmtIdx == -1 { //there is no comment
			line = val
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

func parseLine(line string) (string, ConfVal) {
	endTitle := strings.Index(line, "@")
	var confVal ConfVal
	if endTitle == -1 { // no title (which is okay)
		confVal.title = "Aria's notes" // TODO: make file specific default
	} else { // found title
		confVal.title = strings.TrimSpace(line[:endTitle])
		line = line[endTitle + 1:]
	}
	
	
	splitter := strings.Index(line, "=>")
	if splitter == -1 {
		//error
		return "", confVal
	}
	url   := strings.TrimSpace(line[:splitter])
	files := strings.Split(line[splitter + 2:], ",")
	confVal.files = make([]string, len(files))
	// we need error checking to have non-zero lengths?
	for idx, val := range(files) {
		confVal.files[idx] = strings.TrimSpace(val)
	}
	return url, confVal
}
