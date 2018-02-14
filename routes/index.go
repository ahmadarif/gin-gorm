package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Hello world!"})
	})

	v1 := router.Group("/api/v1/todos")
	v1.POST("", createTodo)
	v1.GET("", fetchAllTodo)
	v1.GET("/:id", fetchSingleTodo)
	v1.PUT("/:id", updateTodo)
	v1.DELETE("/:id", deleteTodo)

	router.Run()
}
