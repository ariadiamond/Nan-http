package main;

import (
    "io/ioutil"
    "strings"
)

const ( 
	NEVER    = 0x1
	SUREAD   = 0x2
	SUWRITE  = 0x4
	READONLY = 0x8
)

func CreateACL () {
	ACL = make(map[string]int)
	contents, err := ioutil.ReadFile(".httpacl")
	if err != nil {
		return
	}
	ACL[".httpacl"] = NEVER
	state := NEVER // start in a safe state
	list  := strings.Split(string(contents), "\n")

	for _, val := range(list) {
        // Remove comments
        line   := val
		cmtIdx := strings.Index(val, "#")
		if cmtIdx >= 0 { // cut out comment. If cmtIdx < 0, there is no comment
			line = val[:cmtIdx]
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
        	state = NEVER
        	if strings.Contains(line, "never") {
        		state = NEVER
        		continue //Skip all other things
        	}
        	if strings.Contains(line, "suRead") {
        		state |= SUREAD
        		state &= ^NEVER // negation
        	}
        	if strings.Contains(line, "suWrite") {
        		state |= SUWRITE
        		state &= ^NEVER
        	}
        	if strings.Contains(line, "readOnly") {
        		state |= READONLY
        		state &= ^NEVER
        	}
        }
        
	}

	//Print the list
	forbiddens := ""
	suRead     := ""
	suWrite    := ""
	readOnly   := ""
	for idx, val := range(ACL) {
        if (val & NEVER) > 0 {
		    forbiddens += idx + " "
		}
		if (val & SUREAD) > 0 {
		    suRead += idx + " "
		}
		if (val & SUWRITE) > 0 {
		    suWrite += idx + " "
		}
		if (val & READONLY) > 0 {
		    readOnly += idx + " "
		}
	}
	// Print values
	Warn("Access Control Protections")
	Warn("Never: " + forbiddens)
	Warn("Sudo Read: " + suRead)
	Warn("Sudo Write: " + suWrite)
	Warn("Read Only: " + readOnly)
}

