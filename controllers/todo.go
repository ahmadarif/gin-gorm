package controllers

import (
	. "ahmadarif/gin-gorm/config"
	"ahmadarif/gin-gorm/models"
	"ahmadarif/gin-gorm/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TodoInsert(c *gin.Context) {
	completed, _ := strconv.ParseBool(c.PostForm("completed"))
	todo := models.Todo{Title: c.PostForm("title"), Completed: completed}
	DB.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

func TodoGetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	//pagination 1 : menggunakan reflect
	var todos []models.Todo
	DB.Find(&todos)
	p := utils.Paginate(todos, page, limit)
	c.JSON(http.StatusOK, p)

	// pagination 2: menggunakan query limit offset
	//var todos []models.Todo
	//var total int
	//offset := utils.CalcOffset(limit, page)
	//DB.Find(&todos).Count(&total)
	//DB.Offset(offset).Limit(limit).Find(&todos)
	//c.JSON(http.StatusOK, gin.H{
	//	"status": http.StatusOK,
	//	"data":   todos,
	//	"total":  total,
	//})
}

func TodoGetByID(c *gin.Context) {
	var todo models.Todo
	DB.First(&todo, c.Param("id"))

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}

func TodoUpdate(c *gin.Context) {
	var todo models.Todo
	DB.First(&todo, c.Param("id"))

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	DB.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.ParseBool(c.PostForm("completed"))
	DB.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

func TodoDelete(c *gin.Context) {
	var todo models.Todo
	DB.First(&todo, c.Param("id"))

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
