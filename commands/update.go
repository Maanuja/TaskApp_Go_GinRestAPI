package commands

import (
	"context"
	"encoding/json"

	"github.com/urfave/cli/v2"

	"github.com/EarvinKayonga/tasks/database"
	"github.com/EarvinKayonga/tasks/models"
)

func UpdateTask(c *cli.Context) error {
	store, err := database.NewJSONFile(c.String("json_file"))
	if err != nil {
		return err
	}

	task, err := deserialize(c.Args().First())
	if err != nil {
		return err
	}

	task, err = store.UpdateTask(context.Background(), *task)
	if err != nil {
		return err
	}

	PrintTasks(*task)

	return nil
}

func deserialize(taskStr string) (*models.Task, error) {
	var task models.Task

	err := json.Unmarshal([]byte(taskStr), &task)
	if err != nil {
		return nil, err
	}

	return models.CreateTask(
		task.ID,
		task.Title,
		task.Status,
		task.CreatedAt,
	), nil
}
