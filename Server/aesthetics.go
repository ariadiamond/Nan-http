package main

import (
    "fmt"
    "os"
)

// Colors
// these do not need to be global among the program, which is why they are lowercase
const (
    red     = "\x1b[31m"
    green   = "\x1b[92m"
    yellow  = "\x1b[33m"
    blue    = "\x1b[34m"
    magenta = "\x1b[35m"
    cyan    = "\x1b[36m"
    unset   = "\x1b[0m"
)

/* Error writes an error to stderr. This prints the message regardless of the Verbosity mode. */
func Error (str string) {
    fmt.Fprintf(os.Stderr, "[%sERR%s] : %s\n", red, unset, str)
}

/* Warn prints things that do not break functionality of the server, but are probably not the
 * intended behaviors. This is used when parsing configuration files, as well as printing important
 * messages that are informational (such as shutting down the server).
 */
func Warn (str string) {
    if Verbosity >= 1 {
        fmt.Printf("[%sWARN%s]: %s\n", blue, unset, str)
    }
}

/* Info prints when an access is made (or at least an attempt). It prints the type of each request
 * in a different color for easy differentiation, along with the path of the object being requested.
 */
func Info (op string, file string) {
    if Verbosity >= 2 {
        var print string
        switch op {
        case "GET":
            print = magenta + "[=] GET   "
        case "POST":
            print = yellow + "[+] POST  "
        case "PUT":
            print = cyan + "[*] PUT   "
        case "DELETE":
            print = red + "[-] DELETE"
        default:
            print = cyan + op
        }
        fmt.Printf("[%sINFO%s]: %s %s%s\n", blue, unset, print, file, unset)
    }
}

/* Usage prints when an invalid command line argument is included. It could be that the option is
 * not recognized, or a valid port number is not included. It takes the binary as the argument to
 * allow for flexibility with binary names (not just Nan)
 */
func Usage (arg string) {
    fmt.Fprintf(os.Stderr, "Usage: %s [-v|V] [-pc] port\n", arg)
    fmt.Fprintf(os.Stderr, "\t-c disable caching of constructed pages\n")
    fmt.Fprintf(os.Stderr, "\t-i run server as HTTP and not HTTPS\n")
    fmt.Fprintf(os.Stderr, "\t-p allow PUT requests\n")
    fmt.Fprintf(os.Stderr, "\t-v verbose\n")
    fmt.Fprintf(os.Stderr, "\t-V very verbose\n")
    fmt.Fprintf(os.Stderr, "\t\x1b[4mport\x1b[0m port to run the server on\n")
    os.Exit(2)
}

/* Start pretty prints a bunch of information when starting an instance of the server. It prints the
 * port the server is running on, whether it is running using HTTP or HTTPs, as well as access
 * control information. I do not plan to keep the current implementation of access control, as it is
 * poorly designed and does not make much sense.
 */
func Start (port int, insecure bool) {
    // This is a friendly server, just like me :)
    fmt.Printf("%sGood Morning!%s\n", green, unset)

    fmt.Printf("%sStarting server on port %d\n", cyan, port)
    fmt.Printf("Verbosity mode: ")
    switch (Verbosity) {
    case 0:
        fmt.Printf("Errors only\n")
    case 1:
        fmt.Printf("Errors and warnings\n")
    case 2:
        fmt.Printf("Every endpoint hit\n")
    default:
        fmt.Printf("Unrecognized%s\n", unset)
        os.Exit(2)
    }
    if insecure {
        fmt.Printf("%sRunning as HTTP and not HTTPS%s\n", red, cyan)
    }
    fmt.Printf("Caching: ")
    if Cache {
        fmt.Printf("%senabled%s\n", green, unset)
    } else {
        fmt.Printf("%sdisabled%s\n", red, unset)
    }
}

/* End just prints good night for when the server is shutting down. Nan is a very friendly server :)
 */
func End () {
    fmt.Printf("\n%sgood night%s\n", green, unset)
}
