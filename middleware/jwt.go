package middleware

import (
	"gin_vue_blog/utils"
	"gin_vue_blog/utils/errormsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour * 10)
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin_vue_blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", errormsg.ERROR
	}
	return tokenString, errormsg.SUCCESS
}

// 验证token
func CheckToken(tokenString string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := setToken.Claims.(*MyClaims); code && setToken.Valid {
		return key, errormsg.SUCCESS
	} else {
		return nil, errormsg.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := errormsg.SUCCESS
		if tokenHeader == "" {
			code = errormsg.ErrorUserTokenNotExists
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errormsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}

		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errormsg.ErrorUserTokenTypeWrong
			c.Abort()
		}
		key, Tcode := CheckToken(checkToken[1])
		if Tcode == errormsg.ERROR {
			code = errormsg.ErrorUserTokenWrong
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errormsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errormsg.ErrorUserTokenExpired
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errormsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
