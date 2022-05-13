package router

import (
	"fmt"
	"logoinig/internal/security/jwt"

	"github.com/gin-gonic/gin"
)

func Run(endPoint string) error {
	r := gin.New()
	r.GET("/admin", Admin)
	r.POST("/login", Login)

	return r.Run(endPoint)
}
func Admin(ctx *gin.Context) {
	token := ctx.Query("token")
	loginInfo, ok := jwt.IsLogedin(token)
	if !ok {
		ctx.String(401, "Invaild token bro ❌")
		return
	}
	ctx.String(200, fmt.Sprintf("hello %s.\n", loginInfo.Username))
}
func Login(ctx *gin.Context) {
	usr, pass := ctx.PostForm("username"), ctx.PostForm("password")
	token, ok := jwt.SingIn(jwt.LoginInfo{Username: usr, Password: pass})
	if !ok {
		ctx.String(400, "false")
		return
	}
	ctx.String(200, "token: "+token)
}
