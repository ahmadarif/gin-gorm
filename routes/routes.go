package routes

import (
	"ahmadarif/gin-gorm/controllers"
	"ahmadarif/gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		// must have headers -> token:mytoken
		sample := v1.Group("/sample", middleware.FakeAuth())
		{
			sample.GET("", controllers.SampleIndex)
			sample.POST("/postForm", controllers.SamplePostForm)
			sample.GET("/query", controllers.SampleQueryString)
			sample.POST("/upload", controllers.SampleUploadFile)
		}

		todos := v1.Group("/todos")
		{
			todos.GET("", controllers.TodoGetAll)
			todos.GET("/db", controllers.TodoGetAllDb)
			todos.GET("/byId/:id", controllers.TodoGetByID)
			todos.POST("", controllers.TodoInsert)
			todos.PUT("/:id", controllers.TodoUpdate)
			todos.DELETE("/:id", controllers.TodoDelete)
		}
	}

	router.Run()
}
