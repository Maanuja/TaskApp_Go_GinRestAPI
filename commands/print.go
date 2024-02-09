package commands

import (
	"encoding/json"
	"fmt"

	"github.com/EarvinKayonga/tasks/models"
)

func PrintTasks(tasks ...models.Task) {
	numberOfTasks := len(tasks)

	if numberOfTasks == 0 {
		fmt.Println("no tasks to print")
		return
	}

	fmt.Printf("printing %d tasks\n", numberOfTasks)
	for _, task := range tasks {
		data, err := jsonify(task)
		if err == nil {
			fmt.Println(data)
		}
	}
}

func jsonify(data models.Task) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
