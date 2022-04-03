package model

import (
	"encoding/base64"
	"gin_vue_blog/utils/errormsg"
	"log"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(id uint, username string) int {
	var user User
	db.Where("username = ?", username).First(&user)
	log.Print("user.ID = ", user.ID, " id= ", id)
	if user.ID > 0 && user.ID != id {
		return errormsg.ErrorUserExists
	}
	return errormsg.SUCCESS
}

// 新增用户
func AddUser(user *User) int {
	user.Password = ScyptPassword(user.Password)
	err := db.Create(user).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户信息
func UpdateUser(id uint, user *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role

	err := db.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 密码加密
func ScyptPassword(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 34, 56, 78, 90, 12, 34, 56}
	HashPw, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
