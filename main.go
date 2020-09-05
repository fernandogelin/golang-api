package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/fernandogelin/golang-api/models" 
	"github.com/fernandogelin/golang-api/controllers" 
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	models.ConnectDatabase()

  r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}