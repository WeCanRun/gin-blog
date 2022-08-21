package model

import (
	"context"
	otgorm "github.com/EDDYCJY/opentracing-gorm"
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

func GetArticles(ctx context.Context, pageNum, pageSize uint) (articles []Article, err error) {
	err = otgorm.WithContext(ctx, db).Select("*").
		Where("state = 1").
		Offset(pageNum).Limit(pageSize).Find(&articles).Error
	return
}

func GetArticleById(ctx context.Context, id uint) (article Article, err error) {
	err = otgorm.WithContext(ctx, db).Select("*").Where("id = ?", id).First(&article).Error
	return
}

// 根据 maps 获取文章总数
func GetArticleTotal(ctx context.Context, maps Article) (count uint) {
	otgorm.WithContext(ctx, db).Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticleByTagId(ctx context.Context, tagId uint, state int) (article []Article, err error) {
	err = otgorm.WithContext(ctx, db).Select("*").Where("tag_id = ? and state = ?", tagId, state).Find(&article).Error
	return
}

func GetArticleByTitle(ctx context.Context, title string, state int) (article []Article, err error) {
	err = otgorm.WithContext(ctx, db).Select("*").Where("title = ? and state = ?", title, state).Find(&article).Error
	return
}

func ExitArticleWithTitle(ctx context.Context, name string) bool {
	var article Article
	otgorm.WithContext(ctx, db).Select("id").Where("title = ?", name).First(&article)
	return article.ID > 0
}

func AddArticle(ctx context.Context, article Article) error {
	return otgorm.WithContext(ctx, db).Create(&article).Error
}

func DeleteArticle(ctx context.Context, id uint) error {
	return otgorm.WithContext(ctx, db).Model(Article{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_at": time.Now(),
		"state":      0,
	}).Error
}

func EditArticle(ctx context.Context, article Article) error {
	return otgorm.WithContext(ctx, db).Model(article).Updates(&article).Where("id = ?", article.ID).Error
}
