package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.GET("/:name/*action", getting)

	// 限制表单上传大小 8MB
	r.MaxMultipartMemory = 8 << 20

	r.GET("/welcome", getting)
	r.POST("/form", posting_form)
	r.POST("/upload", posting_file)
	r.Run(":1111") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getting(c *gin.Context) {
	/*
		c.JSON(200, gin.H{
			"message": "pong",
		})
		c.String(http.StatusOK, "hello world")
	*/
	/*
		// /:name/*action
			name := c.Param("name")
			action := c.Param("action")
			c.String(http.StatusOK, name+" is"+action)
	*/
	/**/
	// 给一个默认值
	name := c.DefaultQuery("name", "hunt")
	c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
}

func posting_form(c *gin.Context) {
	// 表单设置默认值
	type1 := c.DefaultPostForm("type", "alert")
	// 接收其他的
	username := c.PostForm("username")
	passwd := c.PostForm("passwd")
	// 多选框
	hobby := c.PostFormArray("hobby")

	c.String(http.StatusOK, fmt.Sprintf("type: %s, username: %s, passwd: %s, hobby: %v", type1, username, passwd, hobby))
}

func posting_file(c *gin.Context) {

	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	// 传到项目根目录
	c.SaveUploadedFile(file, file.Filename)
	// 打印信息
	c.String(http.StatusOK, fmt.Sprintf("'%s' upload!", file.Filename))
}
