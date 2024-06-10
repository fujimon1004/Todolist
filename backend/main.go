// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "net/http"
)

type Todo struct {
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Status bool   `json:"status"`
}

var db *gorm.DB

func main() {
    var err error
    db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    db.AutoMigrate(&Todo{})

    r := gin.Default()
    r.GET("/todos", getTodos)
    r.POST("/todos", createTodo)
    r.PUT("/todos/:id", updateTodo)
    r.DELETE("/todos/:id", deleteTodo)

    r.Run(":8080")
}

func getTodos(c *gin.Context) {
    var todos []Todo
    db.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
    var todo Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&todo)
    c.JSON(http.StatusOK, todo)
}

func updateTodo(c *gin.Context) {
    id := c.Param("id")
    var todo Todo
    if err := db.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

func deleteTodo(c *gin.Context) {
    id := c.Param("id")
    var todo Todo
    if err := db.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    db.Delete(&todo)
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
