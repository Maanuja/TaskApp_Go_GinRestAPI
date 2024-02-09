package commands

import (
	"context"

	"github.com/urfave/cli/v2"

	"github.com/EarvinKayonga/tasks/database"
)

func CreateTask(c *cli.Context) error {
	store, err := database.NewJSONFile(c.String("json_file"))
	if err != nil {
		return err
	}

	task, err := deserialize(c.Args().First())
	if err != nil {
		return err
	}

	task, err = store.CreateTask(context.Background(), *task)
	if err != nil {
		return err
	}

	PrintTasks(*task)

	return nil
}
