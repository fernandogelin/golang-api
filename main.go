package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/fernandogelin/golang-api/models" 
	"github.com/fernandogelin/golang-api/controllers" 
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

  r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

  r.GET("/books/:id", controllers.FindBook) 
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id")
	r.Run()
}