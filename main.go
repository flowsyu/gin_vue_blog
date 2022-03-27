package main

import (
	"gin_vue_blog/modle"
	"gin_vue_blog/routers"
)

func main() {
	// 引用数据库
	modle.InitDB()

	routers.InitRouter()
}
