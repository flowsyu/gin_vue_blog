package v1

import (
	"gin_vue_blog/server"
	"gin_vue_blog/utils/errormsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size

	url, code := server.UploadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errormsg.GetErrorMsg(code),
		"data": url,
	})
}
