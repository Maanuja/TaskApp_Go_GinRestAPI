package commands

import (
	"github.com/urfave/cli/v2"
)

func Create() *cli.App {
	return &cli.App{
		Name: "tasks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "json_file",
				Aliases: []string{"f"},
				Value:   "tasks.json",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List all tasks",
				Action: func(c *cli.Context) error {
					return ListAll(c)
				},
			},
			{
				Name:  "get",
				Usage: "Get a task by ID",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Value: "",
					},
				},
				Action: func(c *cli.Context) error {
					return GetTaskByID(c)
				},
			},
			{
				Name:  "create",
				Usage: "Create a new task",
				Action: func(c *cli.Context) error {
					return CreateTask(c)
				},
			},
			{
				Name:  "update",
				Usage: "Update a task",
				Action: func(c *cli.Context) error {
					return UpdateTask(c)
				},
			},
			{
				Name:  "delete",
				Usage: "Delete a task by ID",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Value: "",
					},
				},
				Action: func(c *cli.Context) error {
					return DeleteTaskByID(c)
				},
			},
			{
				Name:  "serve",
				Usage: "Start the server",
				// Flags: []cli.Flag{
				// 	&cli.StringFlag{
				// 		Name:    "port",
				// 		Aliases: []string{"pt"},
				// 		Value:   "8083",
				// 	},
				// 	&cli.StringFlag{
				// 		Name:    "host",
				// 		Aliases: []string{"ht"},
				// 		Value:   "localhost",
				// 	},
				// },
				Action: func(c *cli.Context) error {
					return httpServer(c)
				},
			},
		},
	}
}
