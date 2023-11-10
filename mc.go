package main

import (
    "fmt"
    "os"
    "os/exec"
)

const program_name string = "mc"

func usage() {
    fmt.Println(program_name + " <command> ...")
}

func main() {
    if len(os.Args[1:]) == 0 {
        usage()
        os.Exit(0)
    }

    path, err := exec.LookPath(program_name + "-" + os.Args[1])
    if err != nil {
        fmt.Println("Could not find subcommand " + os.Args[1])
        os.Exit(1)
    }

    fmt.Println(path)
}

