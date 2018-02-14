package routes

import (
	"net/http"
	"strconv"

	. "ahmadarif/gin-gorm/config"

	"ahmadarif/gin-gorm/json"
	"ahmadarif/gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func createTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := models.Todo{Title: c.PostForm("title"), Completed: completed}
	DB.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

func fetchAllTodo(c *gin.Context) {
	var todos []models.Todo
	var data []json.Todo

	DB.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	// transforms the todos for building a good response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		}
		data = append(data, json.Todo{ID: item.ID, Title: item.Title, Completed: completed})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func fetchSingleTodo(c *gin.Context) {
	var todo models.Todo
	todoID := c.Param("id")

	DB.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	}

	data := json.Todo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func updateTodo(c *gin.Context) {
	var todo models.Todo
	todoID := c.Param("id")

	DB.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	DB.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	DB.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

func deleteTodo(c *gin.Context) {
	var todo models.Todo
	todoID := c.Param("id")

	DB.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
