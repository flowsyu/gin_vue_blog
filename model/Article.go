package model

import (
	"gin_vue_blog/utils/errormsg"
	"github.com/jinzhu/gorm"
	"log"
)

type Article struct {
	gorm.Model
	Category Category `gorm:"foreignkey:Cid"`
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Img      string   `gorm:"type:varchar(100)" json:"img"`
}

// 新增文章
func AddArticle(article *Article) int {
	err := db.Create(&article).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 查询分类下的所有文章
func GetCateArt(id uint, pageSize, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Where("cid = ?", id).Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&articles).Error
	if err != nil {
		return nil, errormsg.ErrorCategoryNotExists
	}
	return articles, errormsg.SUCCESS
}

// 查询单个文章
func GetArtInfo(id uint) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	log.Println(article, err)
	if err != nil {
		return article, errormsg.ErrorArticleNotExists
	}
	return article, errormsg.SUCCESS
}

// 查询文章列表
func GetArticles(pageSize, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Offset(pageSize * (pageNum - 1)).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errormsg.ERROR
	}
	return articles, errormsg.SUCCESS
}

// 编辑文章
func UpdateArticle(id uint, article *Article) int {
	maps := make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img
	err := db.Model(&Article{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// 删除文章
func DeleteArticle(id uint) int {
	err := db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
