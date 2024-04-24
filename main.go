package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default() //携带基础中间件启动
	router.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "celeste")
		pwd := c.Query("pwd")
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})
	router.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "celeste")
		pwd := c.PostForm("pwd")
		c.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	router.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})
	router.PUT("/path", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
	router.Run(":8080")
}
