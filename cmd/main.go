package main

import (
	"fmt"
	"os"

	"github.com/lucasvmiguel/task/internal/calculator/command"
)

// the format of the cli will be:
// task <ACTION> <TASK_ID>
//
// eg:
// task start task-123
func main() {
	// build a struct with the command written by the user
	cmd, err := command.New(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// a command will run depending of the action that the user has passed
	switch cmd.Action {
	case "start":
		command.Start()
	default:
		fmt.Println("invalid command")
		os.Exit(1)
	}
}
