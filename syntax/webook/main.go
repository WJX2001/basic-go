package main

import (
	"github.com/WJX2001/basic-go/syntax/webook/internal/web/user"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	u := &user.UserHandler{}
	u.RegisterRoutesUser(server)
	server.Run("localhost:8080")
}
