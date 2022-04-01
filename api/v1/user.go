package v1

import (
	"gin_vue_blog/model"
	"gin_vue_blog/utils/errormsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加用户
func AddUser(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	code := model.CheckUser(user.ID, user.Username)
	if code == errormsg.SUCCESS {
		model.CreateUser(&user)
	} else {
		context.Abort()
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
		"data": user,
	})
}

// 查询用户列表
func GetUsers(context *gin.Context) {
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	pageNum, _ := strconv.Atoi(context.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetUsers(pageSize, pageNum)
	code := errormsg.SUCCESS
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
		"data": data,
	})
}

// 编辑用户
func UpdateUser(context *gin.Context) {
	var user model.User
	_ = context.ShouldBindJSON(&user)
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.CheckUser(uint(id), user.Username)
	if code == errormsg.SUCCESS {
		model.UpdateUser(uint(id), &user)
	} else {
		context.Abort()
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
	})
}

// 删除用户(软删除)
func DeleteUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.DeleteUser(id)
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
	})
}
