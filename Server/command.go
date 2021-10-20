package main

import (
    "fmt"
    "net/http"
    "os"
    "regexp"
)

const (
    BUFF_SIZE = 4096
)


func ReadCmd (srv *http.Server) {
    for ;; { // make this an infinite loop to read any commands that come in
        buff := make([]byte, BUFF_SIZE)
        _, err := os.Stdin.Read(buff)
        if err != nil {
            fmt.Print("Error!\n")
        }
        exitRE := regexp.MustCompile(`(\s)*?(?i:e|exit)(\s)+?`)
        if exitRE.FindIndex(buff) != nil {
            // exit
            process, _ := os.FindProcess(os.Getpid())
            process.Signal(os.Interrupt)
            return
        }
        resetRE := regexp.MustCompile(`(\s)*?(?i:r|reset)(\s)+?`)
        if resetRE.FindIndex(buff) != nil {
            Config = make(map[string](map[string]ConfVal))
            Warn("Reset Config")
        }
    }
}
