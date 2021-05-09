package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var Config map[string](map[string][]string)

func ReadConfig(url string) ([]string, bool) {
	lastIndex := strings.LastIndex(url, "/")
	folder    := url[:lastIndex + 1]
	file      := url[lastIndex + 1:]
	folderConfig, exist := Config[folder]
	if !exist { // If we don't already have it, try to get it
		if !ParseConfig(folder) {
			return nil, false
		}
		folderConfig, _ = Config[folder]
	}
	
	files, exist := folderConfig[file]
	if !exist || files == nil {
		return nil, false
	}
	return files, true
}

func ParseConfig(folder string) (bool) {
	contents, err := ioutil.ReadFile(folder + ".httpconfig")
	fmt.Println("folder: ", folder)
	if err != nil { // so we don't parse it again
		Config[folder] = make(map[string][]string)
		return false
	}

	list       := strings.Split(string(contents), "\n")
	fileConfig := make(map[string][]string)
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
		if files != nil {
			fmt.Println("url: ", url, " | files: ", files)
		} else {
			fmt.Println("url: ", url, " | files is nil")
		}
		fileConfig[url] = files
	}
	Config[folder] = fileConfig
	return true
}

func parseLine(line string) (string, []string) {
	splitter := strings.Index(line, "=>")
	if splitter == -1 {
		//error
		return "", nil
	}
	url   := strings.TrimSpace(line[:splitter])
	files := strings.Split(line[splitter + 2:], ",")
	// we need error checking to have non-zero lengths?
	for idx, val := range(files) {
		files[idx] = strings.TrimSpace(val)
	}
	return url, files
}
