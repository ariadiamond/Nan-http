package main;

import (
    "strings"
    "io/ioutil"
)

const ( 
	NEVER  = 0x1
	SUREAD = 0x2
	SUWRITE = 0x4
	READONLY = 0x8
)

func CreateACL () {
	ACL = make(map[string]int)
	contents, err := ioutil.ReadFile(".httpacl")
	if err != nil {
		Warn("Missing .httpignore, did you want one?")
		return
	}
	ACL[".httpignore"] = NEVER
	state := NEVER // start in a safe state
	list := strings.Split(string(contents), "\n")

	// check for comments (#)
	for _, val := range(list) {
        // Remove comments
        line := val
		cmtIdx := strings.Index(val, "#")
		if cmtIdx > 0 {
			line = val[:cmtIdx]
		} else if cmtIdx == 0 { // there is no comment
			line = ""
		}
        // If we don't have any data on here
		if len(line) == 0 {
			continue
		}

        // Do we have a metacharacter
        metaIndex := strings.Index(line, "@")
        if metaIndex < 0 { // no metacharacter, add to ACL
        	ACL[line] = state
        } else {
        	state = 0
        	if strings.Contains(line, "never") {
        		state = NEVER
        	}
        	if strings.Contains(line, "suRead") {
        		state |= SUREAD
        	}
        	if strings.Contains(line, "suWrite") {
        		state |= SUWRITE
        	}
        	if strings.Contains(line, "readOnly") {
        		state |= READONLY
        	}
        }
        
	}

	//Print the list
	forbiddens := ""
	for idx, _ := range(ACL) {
		forbiddens += idx + " "
	}
	Warn("Forbidden list: " + forbiddens)
}

