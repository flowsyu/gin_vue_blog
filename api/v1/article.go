package v1

import (
	modle "gin_vue_blog/model"
	"gin_vue_blog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"strconv"
)

// 添加文章
func AddArticle(context *gin.Context) {
	var article modle.Article
	_ = context.ShouldBindJSON(&article)

	code := modle.AddArticle(&article)
	context.JSON(code, gin.H{
		"code":    code,
		"data":    article,
		"message": errormsg.GetErrorMsg(code),
	})

}

// 获取文章列表
func GetArticles(context *gin.Context) {
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	pageNum, _ := strconv.Atoi(context.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	articles, code := modle.GetArticles(pageSize, pageNum)
	context.JSON(code, gin.H{
		"code":    code,
		"data":    articles,
		"message": errormsg.GetErrorMsg(code),
	})
}

// 查询分类下的所有文章
func GetCateArt(context *gin.Context) {
	cateId, _ := strconv.Atoi(context.Query("id"))
	pageSize, _ := strconv.Atoi(context.Query("pageSize"))
	pageNum, _ := strconv.Atoi(context.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	articles, code := modle.GetCateArt(uint(cateId), pageSize, pageNum)
	context.JSON(code, gin.H{
		"code":    code,
		"data":    articles,
		"message": errormsg.GetErrorMsg(code),
	})
}

// 查询单个文章信息
func GetArtInfo(context *gin.Context) {
	id, _ := strconv.Atoi(context.Query("id"))
	article, code := modle.GetArtInfo(uint(id))
	context.JSON(code, gin.H{
		"code":    code,
		"data":    article,
		"message": errormsg.GetErrorMsg(code),
	})
}

// 编辑文章
func UpdateArticle(context *gin.Context) {
	var article modle.Article
	id, _ := strconv.Atoi(context.Param("id"))
	_ = context.ShouldBindJSON(&article)

	code := modle.UpdateArticle(uint(id), &article)
	context.JSON(code, gin.H{
		"code":    code,
		"data":    article,
		"message": errormsg.GetErrorMsg(code),
	})
}

// 删除文章
func DeleteArticle(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	code := modle.DeleteArticle(uint(id))
	context.JSON(code, gin.H{
		"code":    code,
		"message": errormsg.GetErrorMsg(code),
	})
}
