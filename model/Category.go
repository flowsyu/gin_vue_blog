package model

import (
	"gin_vue_blog/utils/errormsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 检查分类是否存在
func CheckCategory(id uint, name string) int {
	var category Category
	db.Where("name = ?", name).First(&category)
	if category.ID > 0 && category.ID != id {
		return errormsg.ErrorCategoryExists
	}
	return errormsg.SUCCESS
}

// 新增分类
func CreateCategory(category *Category) int {
	err := db.Create(&category).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 查询分类列表
func GetCategories(pageSize, pageNum int) []Category {
	var categories []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categories
}

// 编辑分类信息
func UpdateCategory(id uint, category *Category) int {
	maps := make(map[string]interface{})
	maps["name"] = category.Name
	err := db.Model(&Category{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 删除分类
func DeleteCategory(id int) int {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
