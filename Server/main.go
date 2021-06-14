package main

import (
	"net/http"
	"log"
	"os"
	"strconv"
	"io/ioutil"
	"strings"
)

// Globals
var Forbidden map[string]bool
var Verbosity int
var Sudo      bool

func forbidden () {
	Forbidden = make(map[string]bool)
	if Sudo { // there is nothing that we need to protect against
		return
	}
	contents, err := ioutil.ReadFile(".httpignore")
	if err != nil {
		Warn("Missing .httpignore, did you want one?")
		return
	}
	Forbidden[".httpignore"] = true
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

	//Print the list
	forbiddens := ""
	for idx, _ := range(Forbidden) {
		forbiddens += idx + " "
	}
	Warn("Forbidden list: " + forbiddens)
}

func parseArgs(args []string) (int) {
	var port int
	var err  error
	if len(args) == 2 {
		port, err = strconv.Atoi(args[1])
	} else if len(args) == 3 {
		if args[1][1] == 'V' {
			Verbosity = 2
		} else if args[1][1] == 's' {
			Sudo = true
		} else {
			Usage(args[0])
		}
		port, err = strconv.Atoi(args[2])
	} else if len(args) == 4 {
		if ((args[1][1] != 'v') && (args[2][1] != 'v')) || ((args[1][1] != 's') && (args[2][1] != 's')) {
			Usage(args[0])
		}
		Verbosity = 1
		Sudo      = true
		port, err = strconv.Atoi(args[3])
	} else {
		Usage(args[0])
	}

	if err != nil {
		Usage(args[0])
	}

	return port
}

func main() {
	// Starting up and parsing CLIs
	port := parseArgs(os.Args)
	forbidden()
	Config = make(map[string](map[string]ConfVal))
	Start(port)

	// Actual http server
	http.HandleFunc("/", Handle)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
