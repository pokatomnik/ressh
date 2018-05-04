package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	var args = os.Args[1:]
	for true {
		var command = exec.Command("ssh", args...)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin

		var error = command.Run()

		fmt.Println()
		fmt.Println("Command finished, error(s):")
		if error != nil {
			fmt.Println(error)
		} else {
			fmt.Println("No errors reported")
		}
		fmt.Println()

		time.Sleep(2000 * time.Millisecond)
	}
}
