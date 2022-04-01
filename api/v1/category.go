package v1

import (
	"gin_vue_blog/model"
	"gin_vue_blog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(context *gin.Context) {
	var category model.Category
	_ = context.ShouldBindJSON(&category)
	code := model.CheckCategory(category.ID, category.Name)
	if code == errormsg.SUCCESS {
		model.CreateCategory(&category)
	} else {
		context.Abort()
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
		"data": category,
	})
}

// TODO 查询单个分类下的文章

// 查询分类列表
func GetCategories(context *gin.Context) {
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	pageNum, _ := strconv.Atoi(context.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetCategories(pageSize, pageNum)
	code := errormsg.SUCCESS
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
		"data": data,
	})
}

// 编辑分类
func UpdateCategory(context *gin.Context) {
	var category model.Category
	_ = context.ShouldBindJSON(&category)
	id, _ := strconv.Atoi(context.Param("id"))
	code := model.CheckCategory(uint(id), category.Name)
	if code == errormsg.SUCCESS {
		model.UpdateCategory(uint(id), &category)
	} else {
		context.Abort()
	}
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
	})
}

// 删除分类
func DeleteCategory(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	model.DeleteCategory(id)
	context.JSON(http.StatusOK, gin.H{
		"code": errormsg.SUCCESS,
		"msg":  errormsg.GetErrorMsg(errormsg.SUCCESS),
	})
}
