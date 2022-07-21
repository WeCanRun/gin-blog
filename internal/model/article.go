package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	TagId         uint   `json:"tag_id"` // 与之关联的标签
	Title         string `json:"title"`
	Desc          string `json:"desc"`            //文章描述
	Content       string `json:"content"`         //文章内容
	CoverImageUrl string `json:"cover_image_url"` // 上传的图片地址
	CreatedBy     string `json:"created_by"`
	UpdatedBy     string `json:"updated_by"`
	State         int    `json:"state"` // 状态，0 禁用， 1 启用
	gorm.Model
}

func GetArticles(pageNum, pageSize uint) (articles []Article, err error) {
	err = db.Select("*").
		Where("state = 1").
		Offset(pageNum).Limit(pageSize).Find(&articles).Error
	return
}

func GetArticleById(id uint) (article Article, err error) {
	err = db.Select("*").Where("id = ?", id).First(&article).Error
	return
}

// 根据 maps 获取文章总数
func GetArticleTotal(maps Article) (count uint) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticleByTagId(tagId uint, state int) (article []Article, err error) {
	err = db.Select("*").Where("tag_id = ? and state = ?", tagId, state).Find(&article).Error
	return
}

func GetArticleByTitle(title string, state int) (article []Article, err error) {
	err = db.Select("*").Where("title = ? and state = ?", title, state).Find(&article).Error
	return
}

func ExitArticleWithTitle(name string) bool {
	var article Article
	db.Select("id").Where("title = ?", name).First(&article)
	return article.ID > 0
}

func AddArticle(article Article) error {
	return db.Create(&article).Error
}

func DeleteArticle(id uint) error {
	return db.Model(Article{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
		"state":      0,
	}).Error
}

func EditArticle(article Article) error {
	return db.Model(article).Updates(&article).Where("id = ?", article.ID).Error
}
