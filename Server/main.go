package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Globals

var ACL       map[string]int
var Verbosity int
var SuRead    bool
var SuWrite   bool
var AllowPut  bool

func parseArgs(args []string) (int, bool) {
	port := 0
    insecure := false
	var err error
	for i := 1; i < len(args); i++ {
		if (args[i][0] == '-') { // option
			switch args[i][1] {
            case 'i':
                insecure = true
			case 'p':
				AllowPut = true
			case 'r':
				SuRead = true
			case 'w':
				SuWrite = true
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

	return port, insecure
}

func main() {
	// Starting up and parsing CLIs
	port, insecure := parseArgs(os.Args)

	Config = make(map[string](map[string]ConfVal))
	CreateACL()
	Start(port, insecure)

	// Actual http server
	http.HandleFunc("/", Handle)
    if insecure {
        log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
    } else {
	    log.Fatal(http.ListenAndServeTLS(":" + strconv.Itoa(port), "./Server/cert.pem", "./Server/key.pem", nil))
    }
}
