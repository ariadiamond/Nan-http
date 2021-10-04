package main

import (
    "os" // Stdin
    "fmt"
    "strings"
    "net/http"
)

const (
    BUFF_SIZE = 4096
)


func ReadCmd (srv *http.Server) {
    for ;; {
        buff := make([]byte, BUFF_SIZE)
        _, err := os.Stdin.Read(buff)
        if err != nil {
            fmt.Print("Error!\n")
        }
        bStr := string(buff)
        if strings.Contains(bStr, "e") || strings.Contains(bStr, "exit") {
            // exit
            process, _ := os.FindProcess(os.Getpid())
            process.Signal(os.Interrupt)
            return
        }
        if strings.Contains(bStr, "r") || strings.Contains(bStr, "reset") {
            Config = make(map[string](map[string]ConfVal))
            Warn("Reset Config")
    
        }
    }
}
