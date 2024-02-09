package commands

import (
	"context"

	"github.com/EarvinKayonga/tasks/database"
	"github.com/urfave/cli/v2"
)

func GetTaskByID(c *cli.Context) error {
	store, err := database.NewJSONFile(c.String("json_file"))
	if err != nil {
		return err
	}

	taskID := c.String("id")
	task, err := store.GetTaskByID(context.Background(), taskID)
	if err != nil {
		return err
	}

	PrintTasks(*task)

	return nil
}

func ListAll(c *cli.Context) error {
	store, err := database.NewJSONFile(c.String("json_file"))
	if err != nil {
		return err
	}

	tasks, err := store.GetAllTasks(context.Background())
	if err != nil {
		return err
	}

	PrintTasks(tasks...)

	return nil
}
