package main

import (
	"io/ioutil"
	"strings"
)

var Config map[string](map[string][]string)

func ParseConfig(folder string) {
	contents, err := ioutil.ReadFile(folder + ".httpconfig")
	if err != nil {
		return
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
			continue
		}
		fileConfig[url] = files
	}
	Config[folder] = fileConfig
}

func parseLine(line string) (string, []string) {
	splitter := strings.Index(line, "=>")
	if splitter == -1 {
		//error
		return "", nil
	}
	url   := strings.TrimSpace(line[:splitter])
	files := strings.Split(line[splitter + 1:], ",")
	// we need error checking to have non-zero lengths?
	for idx, val := range(files) {
		files[idx] = strings.TrimSpace(val)
	}
	return url, files
}
