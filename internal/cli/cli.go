package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"tomato/internal/app"
	"tomato/internal/repo"
)

// New creates a new cli application
func New() *cli.App {
	c := cli.NewApp()
	c.Name = "tomato"
	c.Version = "1.0.0"
	c.Usage = "some text"

	cfg := initConfig()
	settings := initSettings(cfg.SettingsPath)
	r := repo.New(initDBConnection(cfg.DBPath))
	a := app.New(settings, r)

	commands := []*cli.Command{
		{
			Name:  "load",
			Usage: "Load task list from `FILE`",

			Flags: []cli.Flag{
				// &cli.BoolFlag{
				// 	Name:  "forced",
				// 	Value: false,
				// 	Usage: "tasks will be load even though current task list is not empty",
				// },
			},
			Action: func(c *cli.Context) error {
				if c.NArg() < 1 {
					return fmt.Errorf("%d", c.NArg())
				}

				return a.LoadTasks(c.Args().First(), true)
			},
		},
		{
			Name:  "edit",
			Usage: "Open task list for edit in the default editor",
			Action: func(c *cli.Context) error {
				return a.EditTasks()
			},
		},
		{
			Name:  "start",
			Usage: "Start the next task from the task list",
			Action: func(c *cli.Context) error {
				return a.StartTask()
			},
		},
		{
			Name:  "stop",
			Usage: "Stop timer for the current task",
			Action: func(c *cli.Context) error {
				return a.StopTask()
			},
		},
		{
			Name:  "log",
			Usage: "Print all of the finished tasks",
			Action: func(c *cli.Context) error {
				return a.GetLog()
			},
		},
		{
			Name:  "list",
			Usage: "Print the current task list",
			Action: func(c *cli.Context) error {
				return a.GetCurrentTasks()
			},
		},
	}

	c.Commands = commands

	c.Action = func(c *cli.Context) error {
		return nil
	}

	return c
}
