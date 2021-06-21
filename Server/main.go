package main

import (
	"net/http"
	"log"
	"os"
	"strconv"
)

// Globals

var ACL       map[string]int
var Verbosity int
var Sudo      bool

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
	CreateACL()
	Config = make(map[string](map[string][]string))
	Start(port)

	// Actual http server
	http.HandleFunc("/", Handle)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
