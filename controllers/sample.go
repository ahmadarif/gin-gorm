package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SampleIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Hello world!"})
}

func SamplePostForm(c *gin.Context) {
	name := c.PostForm("name")
	hello := fmt.Sprintf("Hello %s", name)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": hello})
}

func SampleQueryString(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": gin.H{"firstname": firstname, "lastname": lastname}})
}

func SampleUploadFile(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")

	//Upload the file to specific path.
	if err := c.SaveUploadedFile(file, fmt.Sprintf("tmp/%s", file.Filename)); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{"status": http.StatusInternalServerError, "message": "Upload failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Upload successfully"})
}
