package commands

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"context"
	"github.com/EarvinKayonga/tasks/database"
	"github.com/EarvinKayonga/tasks/models"
	"github.com/urfave/cli/v2"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	db, err := database.NewJSONFile("tasks.json")
	if err != nil {
		log.Fatalf("unable to initialize JSON file-based database: %s", err)
	}

	r.GET("/tasks", func(c *gin.Context) {
		tasks, err := db.GetAllTasks(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		task, err := db.GetTaskByID(context.Background(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"task": task})
	})

	r.POST("/task", func(c *gin.Context) {
		var task models.Task
		if err := c.BindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		createdTask, err := db.CreateTask(context.Background(), task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"task": createdTask})
	})

	r.PUT("/task/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTask models.Task
		if err := c.BindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedTask.ID = id
		_, err := db.UpdateTask(context.Background(), updatedTask)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
	})

	r.DELETE("/task/:id", func(c *gin.Context) {
		id := c.Param("id")
		deletedTask, err := db.DeleteTaskByID(context.Background(), id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully", "task": deletedTask})
	})

	return r
}

func httpServer(c *cli.Context) error {
    r := setupRouter()
    // Listen and Server in 0.0.0.0:8083
    r.Run(":8083")
    return nil
}