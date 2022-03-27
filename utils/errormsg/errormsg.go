package errormsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000...  用户模块的错误
	ErrorUserExists         = 1001 // 用户已经存在
	ErrorUserPasswordWrong  = 1002 // 用户密码错误
	ErrorUserNotExists      = 1003 // 用户不存在
	ErrorUserTokenNotExists = 1004 // 用户token不存在
	ErrorUserTokenExpired   = 1005 // 用户token过期
	ErrorUserTokenWrong     = 1006 // 用户token错误
	ErrorUserTokenTypeWrong = 1007 // 用户token类型错误
	// code = 2000...  文章模块的错误

	// code = 3000...  分类模块的错误

)

var codeMsg = map[int]string{
	SUCCESS:                 "success",
	ERROR:                   "fail",
	ErrorUserExists:         "用户已经存在",
	ErrorUserPasswordWrong:  "用户密码错误",
	ErrorUserNotExists:      "用户不存在",
	ErrorUserTokenNotExists: "用户token不存在",
	ErrorUserTokenExpired:   "用户token过期",
	ErrorUserTokenWrong:     "用户token错误",
	ErrorUserTokenTypeWrong: "用户token类型错误",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
