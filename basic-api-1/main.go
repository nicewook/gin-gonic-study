package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestModel struct {
	Id   int    `json:"id" binding:"required`
	Name string `json:"name" binding:"required`
}

func main() {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		// c.String(http.StatusOK, "hello world!")
		c.JSON(http.StatusOK, gin.H{
			"responseData": "hello JSON wonrld",
		})
	})
	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"greetings": fmt.Sprintf("hello %v", name),
		})
	})
	r.POST("/add", func(c *gin.Context) {
		var data TestModel
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("err: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"dataReceived": data,
		})
	})

	r.Run("localhost:8080")
}
