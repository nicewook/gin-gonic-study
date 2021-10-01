package main

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID     int                   `uri:"id"`
	Name   string                `form:"name"`
	Email  string                `form:"email"`
	Avatar *multipart.FileHeader `form:"avatar"`
}

func putUserHandle(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		log.Println("err: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("user: %+v", user)

	if err := c.ShouldBindUri(&user); err != nil {
		log.Println("err: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	log.Printf("user: %+v", user)

	if err := c.SaveUploadedFile(user.Avatar, "assets/"+user.Avatar.Filename); err != nil {
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
	r.PUT("/user/:id", putUserHandle)
	return r
}

func main() {
	newServer().Run()
}
