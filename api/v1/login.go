package v1

import (
	"gin_vue_blog/middleware"
	"gin_vue_blog/model"
	"gin_vue_blog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	var user model.User
	var token string
	var code int
	context.ShouldBindJSON(&user)

	code = model.Login(user.Username, user.Password)

	if code == errormsg.SUCCESS {
		token, code = middleware.SetToken(user.Username)
	}
	context.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   errormsg.GetErrorMsg(code),
		"token": token,
	})
}
