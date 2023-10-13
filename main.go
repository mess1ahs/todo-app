package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mess1ahs/todo-app/models"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	r := gin.Default()

	// API tasks
	router := r.Group("/tasks")
	{
		router.POST("/", postTask)
		router.GET("/", readTasks)
		router.GET("/:id", readTask)
		router.PUT("/", updateTask)
		router.DELETE("/:id", deleteTask)
	}

	// By default it serves on :8080
	r.Run()
}
func postTask(c *gin.Context) {
	var task struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.CreateTask(task.Title)
	if err != nil {
		c.JSON(505, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "A new Record Created!", "id": id})
}

func readTasks(c *gin.Context) {
	tasks, err := models.GetTasks()
	checkErr(err)

	if tasks == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(200, gin.H{"data": tasks})
	}
}

func readTask(c *gin.Context) {
	id := c.Param("id") // Get the task ID from the URL parameter

	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := models.GetTaskByID(taskID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	if task.Id == 0 {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, gin.H{"data": task})
}

func updateTask(c *gin.Context) {

	var updatedTask struct {
		Id   int  `json:"id"`
		Done bool `json:"done"`
	}

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.UpdateTask(updatedTask.Id, updatedTask.Done)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": "Record Updated!"})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id") // Get the task ID from the URL parameter

	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	err = models.DeleteTask(taskID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": "Record Deleted!"})
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
