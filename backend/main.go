package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   string `json:"id"`
	Task string `json:"task"`
}

var todos []Todo

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.BindJSON(&todo); err == nil {
		todos = append(todos, todo)
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.Run(":8000")
}
