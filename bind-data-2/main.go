package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `form:"id" uri:"id" json:"id"`
	Name  string `form:"name" uri:"name" json:"name"`
	Email string `form:"email" uri:"email" json:"email"`
}

func getUserQueryHandle(c *gin.Context) {
	var user User
	if err := c.ShouldBindQuery(&user); err != nil {
		log.Println("err: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("user: %+v", user)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   user,
	})
}

func postUserJSONHandle(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("err: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("user: %+v", user)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   user,
	})
}

func putUserURIHandle(c *gin.Context) {
	var user User
	if err := c.ShouldBindUri(&user); err != nil {
		log.Println("err: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("user: %+v", user)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   user,
	})
}

func putUserURIJSONHandle(c *gin.Context) {
	var user User
	if err := c.ShouldBindUri(&user); err != nil {
		log.Println("err: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("user: %+v", user)

	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("err: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("user: %+v", user)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   user,
	})
}

func newServer() *gin.Engine {
	r := gin.Default()
	r.GET("/user", getUserQueryHandle)
	r.POST("/user", postUserJSONHandle)
	r.PUT("/user/:id/:name/:email", putUserURIHandle)
	r.PUT("/user/:id", putUserURIJSONHandle)

	return r
}

func main() {
	newServer().Run()
}
