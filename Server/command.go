package main

import (
    "fmt"
    "os"
    "regexp"
)

const (
    buff_size = 4096
)


/* ReadCmd reads a command in an infinite loop from stdin and does any actions specified. If there
 * is no action for the command given, the function does not do anything. Currently actions
 * supported are:
 *     - e/exit: gracefully shuts down the server. This is equivalent to Ctrl/Cmd+C.
 *     - r/reset: clears the config state in memory. This is useful to reflect updates to config
 *       files without having to restart the server.
 */
func ReadCmd () {
    for ;; { // make this an infinite loop to read any commands that come in
        buff := make([]byte, buff_size)
        _, err := os.Stdin.Read(buff)
        if err != nil {
            fmt.Print("Error!\n")
        }

        // exitRE: any amount of space, followed by e or exit, followed by more space
        exitRE := regexp.MustCompile(`(\s)*?(?i:e|exit)(\s)+?`)
        if exitRE.FindIndex(buff) != nil {
            // exit
            process, _ := os.FindProcess(os.Getpid())
            process.Signal(os.Interrupt)
            return
        }

        // resetRE: any amount of space, followed by r or reset, followed by more space 
        resetRE := regexp.MustCompile(`(\s)*?(?i:r|reset)(\s)+?`)
        if resetRE.FindIndex(buff) != nil {
            Config = make(map[string](map[string]ConfVal))
            Warn("Reset Config")
        }
    }
}
