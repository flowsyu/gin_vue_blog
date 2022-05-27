package v1

import (
	"gin_vue_blog/model"
	"gin_vue_blog/utils/errormsg"
	"gin_vue_blog/utils/validator"
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
	msg, code := validator.Validate(&user)
	if code != errormsg.SUCCESS {
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
		})
		return
	}

	code = model.CheckUser(user.ID, user.Username)
	if code == errormsg.SUCCESS {
		model.AddUser(&user)
	} else {
		context.Abort()
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
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

	data, total := model.GetUsers(pageSize, pageNum)
	code := errormsg.SUCCESS
	context.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   errormsg.GetErrorMsg(code),
		"total": total,
		"data":  data,
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
