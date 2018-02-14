package routes

import (
	"ahmadarif/gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	api := router.Group("/api", middleware.FakeAuth())
	api.GET("", index)
	api.POST("", hello)
	api.GET("/query", queryString)
	api.POST("/upload", uploadFile)
	api.GET("/fakeAuth", fakeAuth)

	v1 := router.Group("/api/v1/todos")
	v1.POST("", createTodo)
	v1.GET("", fetchAllTodo)
	v1.GET("/:id", fetchSingleTodo)
	v1.PUT("/:id", updateTodo)
	v1.DELETE("/:id", deleteTodo)

	router.Run()
}
