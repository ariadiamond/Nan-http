package main

import (
    "context"
    "errors"
    "log"
    "net/http"
    "os"
    "os/signal"
    "strconv"
)

// Globals
var ACL       map[string]int
var Verbosity int
var SuRead    bool
var SuWrite   bool
var AllowPut  bool

/* parseArgs reads arguments passed in when initializing the program, and mostly sets global
 * variables. It returns the port specified (this is currently still required, although feasibly
 * defaults could be set). It also returns whether to run the server in http or https. This second
 * parameter could be expanded to run on both http and https.
 */
func parseArgs(args []string) (int, bool) {
    port := 0
    insecure := false // use http instead of https
    var err error
    for i := 1; i < len(args); i++ {
        if (args[i][0] == '-') { // option
            for j := 1; j < len(args[i]); j++ { // iterate through however many option characters
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

/* This essentially calls other functions to do everything. */
func main() {
    // Starting up and parsing CLIs
    port, insecure := parseArgs(os.Args)

    // prep
    Config = make(map[string](map[string]ConfVal))
    CreateACL()
    Start(port, insecure)

    // Set up http server
    http.HandleFunc("/", Handle)
    srv := http.Server { Addr: ":" + strconv.Itoa(port) }

    // Read commands and catch SIGINT
    go ReadCmd(&srv)
    go func() { // anonymous inner function, which allows us to use srv as a variable
        sigint := make(chan os.Signal, 1)
        signal.Notify(sigint, os.Interrupt)
        <- sigint
        Warn("Shutting down server")
        err := srv.Shutdown(context.Background())
        if err != nil {
            Error("Error while shutting down server")
        }
    }()

    // run server
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
