package model

import (
	"fmt"
	"gin_vue_blog/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.DbName),
	)

	if err != nil {
		fmt.Println("数据库连接失败,请检查参数是否正确", err)
	}

	// 禁用表名复数
	db.SingularTable(true)

	db.AutoMigrate(&User{}, &Article{}, &Category{})
	// SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns用于设置最大打开的连接数
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime用于设置连接的最大可复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//err = db.Close()
	//if err != nil {
	//	return
	//}
}
