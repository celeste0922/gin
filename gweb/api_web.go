package gweb

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

type PostParams struct {
	Name string `json:"name" uri:"name" form:"name"`
	Age  int    `json:"age" uri:"age" form:"age"`
	Sex  bool   `json:"sex" uri:"sex" form:"sex"`
}

func middel() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("one...before func")
		c.Next()
		fmt.Println("one...after func")
	}
}
func middelTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("two...before func")
		c.Next()
		fmt.Println("two...after func")
	}
}
func WebStart() {
	r := gin.Default() //携带基础中间件启动
	v1 := r.Group("/path").Use(middel(), middelTwo())
	v2 := r.Group("/test")
	v1.GET("get/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "celeste")
		pwd := c.Query("pwd")
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})
	v1.POST("post", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "celeste")
		pwd := c.PostForm("pwd")
		fmt.Println("post=====>", user, pwd)
		c.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	v1.DELETE("delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})
	v1.PUT("put", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})
	//=======================================================================
	//shouldBindJSON
	v2.POST("bind", func(c *gin.Context) {
		var params PostParams
		err := c.ShouldBindJSON(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "failed",
				"msg":    err,
				"data":   gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": params,
			})
		}
	})
	//shouldBindUri
	v2.POST("bind/:name/:age/:sex", func(c *gin.Context) {
		var params PostParams
		err := c.ShouldBindUri(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "failed",
				"msg":    err,
				"data":   gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": params,
			})
		}
	})
	//shouldBindWithQuery
	v2.POST("bind/query", func(c *gin.Context) {
		var params PostParams
		err := c.ShouldBindQuery(&params)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "failed",
				"msg":    err,
				"data":   gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "success",
				"data": params,
			})
		}
	})
	//upload
	v2.POST("upload", func(c *gin.Context) {
		//form, _ := c.MultipartForm()
		//files := form.File["files"] //多文件
		file, _ := c.FormFile("file")
		//name := c.PostForm("name")
		log.Println(file.Filename)
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		io.Copy(out, in)
		defer out.Close()
		//c.SaveUploadedFile(file, "./"+file.Filename)
		//c.JSON(http.StatusOK, gin.H{
		//	"msg":  file,
		//	"name": name,
		//})
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename="+file.Filename))
		c.File("./" + file.Filename)
	})
	r.Run(":8080")
}
