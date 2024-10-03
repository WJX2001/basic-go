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

	// 参数路由
	server.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello 这是参数路由"+name)
	})

	// 通配符路由
	server.GET("/views/*.html", func(ctx *gin.Context) {
		page := ctx.Param(".html")
		ctx.String(http.StatusOK, "hello 这是通配符路由"+page)
	})

	// 查询参数
	server.GET("/order", func(ctx *gin.Context) {
		oid := ctx.Query("id")
		ctx.String(http.StatusOK, "hello 这是查询参数"+oid)
	})

	server.Run("localhost:8080")
}
