package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yanjianyu-nwpu/golang_json_tool"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("1.html")
	//form
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "1.html", gin.H{
			"title": "Main website",
		})
	})

	router.POST("/index", func(c *gin.Context) {
		src := c.PostForm("src") //可设置默认值
		dst := c.PostForm("dst")
		fmt.Println(src)
		n := golang_json_tool.GenCopyCode(src, dst)
		fmt.Println(n)
		c.String(http.StatusOK, src)
	})

	router.Run(":9527")
}
