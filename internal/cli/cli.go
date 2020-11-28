package cli

import (
	"github.com/urfave/cli/v2"
)

// New creates a new cli application
func New() *cli.App {
	c := cli.NewApp()
	c.Name = "tomato"
	c.Version = "1.0.0"
	c.Usage = "some text"

	commands := []*cli.Command{
		{
			Name:  "load",
			Usage: "Load task list from file",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "edit",
			Usage: "Open task list for edit in the default editor",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "start",
			Usage: "Start the next task from the task list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "Stop timer for the current task",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "log",
			Usage: "Print all of the finished tasks",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "Print the current task list",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	return c
}
