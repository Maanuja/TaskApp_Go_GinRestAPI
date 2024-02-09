package commands

import (
	"context"

	"github.com/urfave/cli/v2"

	"github.com/EarvinKayonga/tasks/database"
)

func DeleteTaskByID(c *cli.Context) error {
	store, err := database.NewJSONFile(c.String("json_file"))
	if err != nil {
		return err
	}

	taskID := c.String("id")
	task, err := store.DeleteTaskByID(context.Background(), taskID)
	if err != nil {
		return err
	}

	PrintTasks(*task)

	return nil
}
