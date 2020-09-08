package command

import (
	"fmt"
	"os"
)

// Command is going to be executed by an user
// format:
// task <ACTION> <TASK_ID>
//
// eg:
// task start task-123
type Command struct {
	Action string
	TaskID string
}

// New created a new Command struct
func New(osArgs []string) (Command, error) {
	if len(osArgs) < 2 {
		return Command{}, fmt.Errorf("invalid number of arguments")
	}

	return Command{
		Action: os.Args[1],
		TaskID: os.Args[2],
	}, nil
}
