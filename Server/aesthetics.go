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

func Error (str string) {
	fmt.Fprintf(os.Stderr, "[%sERR%s] : %s\n", red, unset, str)
}

func Warn (str string) {
	if Verbosity >= 1 {
		fmt.Fprintf(os.Stdout, "[%sWARN%s]: %s\n", blue, unset, str)
	}
}

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
		fmt.Fprintf(os.Stdout, "[%sINFO%s]: %s %s%s\n", blue, unset, print, file, unset)
	}
}

func Usage (arg string) {
	fmt.Fprintf(os.Stderr, "Usage: %s [-v|V] [-prw] port\n", arg)
    fmt.Fprintf(os.Stderr, "\t-i run server as HTTP and not HTTPS\n")
	fmt.Fprintf(os.Stderr, "\t-p allow PUT requests\n")
	fmt.Fprintf(os.Stderr, "\t-v verbose\n")
	fmt.Fprintf(os.Stderr, "\t-V very verbose\n")
	fmt.Fprintf(os.Stderr, "\t-r Sudo read\n")
	fmt.Fprintf(os.Stderr, "\t-w Sudo write\n")
	fmt.Fprintf(os.Stderr, "\t\x1b[4mport\x1b[0m port to run the server on\n")
	os.Exit(2)
}

func Start (port int, insecure bool) {
	fmt.Fprintf(os.Stdout, "%sStarting server on port %d\n", cyan, port)
	fmt.Fprintf(os.Stdout, "Verbosity mode: ")
	switch (Verbosity) {
	case 0:
		fmt.Fprintf(os.Stdout, "Errors only\n")
	case 1:
		fmt.Fprintf(os.Stdout, "Errors and warnings\n")
	case 2:
		fmt.Fprintf(os.Stdout, "Every endpoint hit\n")
	default:
		fmt.Fprintf(os.Stdout, "Unrecognized%s\n", unset)
		os.Exit(2)
	}
    if insecure {
        fmt.Fprintf(os.Stdout, "%sRunning as HTTP and not HTTPS%s\n", red, cyan)
    }
	fmt.Fprintf(os.Stdout, "Sudo Read: ")
	if SuRead {
		fmt.Fprintf(os.Stdout, "%senabled, please be careful%s\n", red, cyan)
	} else {
		fmt.Fprintf(os.Stdout, "%sdisabled%s\n", green, cyan)
	}
	fmt.Fprintf(os.Stdout, "Sudo Write: ")
	if SuWrite {
		fmt.Fprintf(os.Stdout, "%senabled, please be more careful%s\n", red, unset)
	} else {
		fmt.Fprintf(os.Stdout, "%sdisabled%s\n", green, unset)
	}
}
