package main

import (
	"net/http"
	"log"
	"os"
	"strconv"
	"errors"
)

// Globals

var ACL       map[string]int
var Verbosity int
var Sudo      bool

func parseArgs(args []string) (int) {
	port := 0
	var err error
	for i := 1; i < len(args); i++ {
		if (args[i][0] == '-') { // option
			switch args[i][1] {
			case 'p': // TODO
				
			case 's':
				Sudo = true
			case 'v':
				Verbosity = 1
			case 'V':
				Verbosity = 2
			default:
				err = errors.New("Invalid arguments")	
			}
		} else { // port
			port, err = strconv.Atoi(args[i])
		}
		// print usage if invalid
		if err != nil {
			Usage(args[0])
		}
	}
	if port == 0 {
		Usage(args[0])
	}

	return port
}

func main() {
	// Starting up and parsing CLIs
	port := parseArgs(os.Args)

	Config = make(map[string](map[string]ConfVal))
	CreateACL()
	Start(port)

	// Actual http server
	http.HandleFunc("/", Handle)
	log.Fatal(http.ListenAndServeTLS(":" + strconv.Itoa(port), "./Server/cert.pem", "./Server/key.pem", nil))
}
