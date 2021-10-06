package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"os/signal"
	"context"
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
			for j := 1; j < len(args[i]); j++ {
				switch args[i][j] {
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
	srv := http.Server { Addr: ":" + strconv.Itoa(port) }
	go ReadCmd(&srv)
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<- sigint
    	Warn("Shutting down server")
    	err := srv.Shutdown(context.Background())
    	if err != nil {
        	Error("Error while shutting down server")
    	}
	}()
	
	var err error
    if insecure {
        err = srv.ListenAndServe()
    } else {
	    err = srv.ListenAndServeTLS("./Server/cert.pem", "./Server/key.pem")
    }
    if err != http.ErrServerClosed {
    	log.Fatal(err)
    }
}
