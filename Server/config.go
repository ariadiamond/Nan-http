package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var Config map[string](map[string][]string)

func ParseConfig(folder string) {
	contents, err := ioutil.ReadFile(folder + ".httpconfig")
	if err != nil {
		return
	}	

	list := strings.Split(string(contents), "\n")

	// check for comments (#)
	for _, val := range(list) {
		if len(val) == 0 {
			continue
		}
		cmtIdx := strings.Index(val, "#")
		if cmtIdx > 0 {
			Forbidden[val[:cmtIdx]] = true
		} else if cmtIdx == -1 { //there is no comment
			Forbidden[val] = true
		}
	}

}

func parseLine(line string) (string, []string) {
	splitter := strings.Contains(line, "=>")
	if splitter == -1 {
		//error
	}
	//strings.TrimSpace

}
