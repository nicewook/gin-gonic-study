package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestModel struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

var V1Users = []TestModel{
	{
		ID:   1,
		Name: "user v1 - one",
	},
	{
		ID:   2,
		Name: "user v1 - two",
	},
}

var V1Products = []TestModel{
	{
		ID:   1,
		Name: "product v1 - one",
	},
	{
		ID:   2,
		Name: "product v1 - two",
	},
}

var V2Users = []TestModel{
	{
		ID:   1,
		Name: "user v2 - one",
	},
	{
		ID:   2,
		Name: "user v2 - two",
	},
}

var V2Products = []TestModel{
	{
		ID:   1,
		Name: "product v2 - one",
	},
	{
		ID:   2,
		Name: "product v2 - two",
	},
}

func newServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		user := v1.Group("user")
		{
			user.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"data": V1Users,
				})
			})

		}
		product := v1.Group("product")
		{
			product.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"data": V1Products,
				})
			})

		}
	}
	v2 := r.Group("v2")
	{
		user := v2.Group("user")
		{
			user.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"data": V2Users,
				})
			})

		}
		product := v2.Group("product")
		{
			product.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"data": V2Products,
				})
			})
		}
	}
	return r
}

func main() {
	newServer().Run()
}
