package main

import (
	"os"
	"fmt"
)

// Colors
// these do not need to be global
const red   = "\x1b[31m"
const green = "\x1b[92m"
const blue  = "\x1b[34m"
const cyan  = "\x1b[36m"
const unset = "\x1b[0m"

func Error (str string) {
	fmt.Fprintf(os.Stderr, "[%sERR%s] : %s\n", red, unset, str)
}

func Warn (str string) {
	if Verbosity >= 1 {
		fmt.Fprintf(os.Stdout, "[%sWARN%s]: %s\n", blue, unset, str)
	}
}

func Usage (arg string) {
	fmt.Fprintf(os.Stderr, "Usage: %s [-v] port\n", arg)
	fmt.Fprintf(os.Stderr, "\t-p allow PUT (todo)\n")
	fmt.Fprintf(os.Stderr, "\t-v verbose (TODO)\n")
	fmt.Fprintf(os.Stderr, "\tport port to run the server on\n")
	os.Exit(2)
}

func Start (port int) {
	fmt.Fprintf(os.Stdout, "%sStarting server on port %d\n", cyan, port)
	fmt.Fprintf(os.Stdout, "Verbosity mode: ")
	switch (Verbosity) {
	case 0:
		fmt.Fprintf(os.Stdout, "Errors only\n")
	case 1:
		fmt.Fprintf(os.Stdout, "Errors and warnings\n")
	default:
		fmt.Fprintf(os.Stdout, "Unrecognized%s\n", unset)
		os.Exit(2)
	}
	fmt.Fprintf(os.Stdout, "Sudo mode: ")
	if Sudo {
		fmt.Fprintf(os.Stdout, "%senabled, please be careful%s\n", red, unset)
	} else {
		fmt.Fprintf(os.Stdout, "%sdisabled%s\n", green, unset)
	}
}
