package main

import (
	modle "gin_vue_blog/model"
	"gin_vue_blog/routers"
)

func main() {
	// 引用数据库
	modle.InitDB()

	routers.InitRouter()
}
