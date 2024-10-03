package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// 当一个 HTTP请求， 用GET方法访问的时候，如果访问路径 /hello
	server.GET("/hello", func(c *gin.Context) {
		// 就执行这段代码
		c.String(http.StatusOK, "hello go")
	})

	server.POST("/helloPost", func(c *gin.Context) {
		c.String(http.StatusOK, "hello go post")
	})

	server.Run("localhost:8080")
}
