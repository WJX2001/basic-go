package user

import "github.com/gin-gonic/gin"

// UserHandler 在此定义跟 user有关的路由
type UserHandler struct {
}

// 这种写法缺陷：容易被别人注册相同的路由
func (u *UserHandler) RegisterRoutesUser(server *gin.Engine) {
	// 统一处理前缀
	ug := server.Group("/user")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	// server.GET("/user/profile", u.Profile)
	// server.POST("/user/signup", u.SignUp)
	// server.POST("/user/login", u.Login)
	// server.POST("/user/edit", u.Edit)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {

}

func (u *UserHandler) Login(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
