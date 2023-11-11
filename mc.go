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

	subcommand_exec := program_name + "-" + os.Args[1]

	_, err := exec.LookPath(subcommand_exec)
	if err != nil {
		fmt.Println("Could not find subcommand " + os.Args[1])
		os.Exit(1)
	}

	cmd := exec.Command(subcommand_exec, os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
