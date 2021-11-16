package main

import (
    "context"
    "errors"
    "net/http"
    "os"
    "os/signal"
    "strconv"
)

// Globals
var ACL       map[string]int
var Verbosity int
var AllowPut  bool

/* parseArgs reads arguments passed in when initializing the program, and mostly sets global
 * variables. It returns the port specified (this is currently still required, although feasibly
 * defaults could be set). It also returns whether to run the server in http or https. This second
 * parameter could be expanded to run on both http and https.
 */
func parseArgs(args []string) (int, bool) {
    port := 0
    insecure := false // use HTTP instead of HTTPS. By default we do not want to use HTTP
    var err error
    for i := 1; i < len(args); i++ { // iterate through command line arguments
        if (args[i][0] == '-') { // option
            for j := 1; j < len(args[i]); j++ { // iterate through however many option characters
                switch args[i][j] {
                case 'i':
                    insecure = true
                case 'p':
                    AllowPut = true
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
    if port == 0 { // port number was not set
        Usage(args[0])
    }

    return port, insecure
}

/* main essentially calls other functions to do everything. */
func main() {
    // Starting up and parsing CLIs
    port, insecure := parseArgs(os.Args)

    // prep
    Config = make(map[string](map[string]ConfVal))
    Start(port, insecure)

    // Set up http server
    http.HandleFunc("/", Handle)
    srv := http.Server { Addr: ":" + strconv.Itoa(port) }

    // Read commands
    go ReadCmd(&srv)
    
    go func() { // anonymous inner function, which allows us to use srv as a variable
        // catch Cmd/Ctrl + C and stop the server gracefully
        sigint := make(chan os.Signal, 1)
        signal.Notify(sigint, os.Interrupt)
        <- sigint // wait on stop before executing further code
        
        End() // This just says good night :)
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
        Error(err.Error())
    }
}
