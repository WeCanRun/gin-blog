package model

import (
	"context"
	otgorm "github.com/EDDYCJY/opentracing-gorm"
	"github.com/jinzhu/gorm"
)

//CREATE TABLE `blog_tag` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`name` varchar(100) DEFAULT '' COMMENT '标签名称',
//`created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
//`created_by` varchar(100) DEFAULT '' COMMENT '创建人',
//`modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
//`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
//`deleted_on` int(10) unsigned DEFAULT '0',
//`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
type Tag struct {
	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	State     int    `json:"state"`
	gorm.Model
}

//func(tag Tag) TableName() string{
//	return "blog_tag"
//}

//func (tag *Tag) BeforeDelete(scope *gorm.Scope) error {
//	scope.SetColumn("DeletedAt", time.Now().Unix())
//	return nil
//}

// 获取所有文章标签
func GetTags(ctx context.Context, pageNum uint, pageSize uint) (tags []Tag, err error) {
	err = otgorm.WithContext(ctx, db).Where("state = ?", 1).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	return
}

func GetTagsByName(ctx context.Context, name string) (tags []Tag, err error) {
	err = otgorm.WithContext(ctx, db).Model(Tag{}).Where("name = ?", name).Find(&tags).Error
	return
}

func DeleteTag(ctx context.Context, id uint) (err error) {
	return otgorm.WithContext(ctx, db).Where("id = ?", id).Delete(Tag{}).Error
}

func AddTag(ctx context.Context, tag Tag) (err error) {
	err = otgorm.WithContext(ctx, db).Model(tag).Create(&tag).Error
	return
}

func EditTag(ctx context.Context, tag Tag) (err error) {
	err = otgorm.WithContext(ctx, db).Model(tag).Where("id = ?", tag.ID).Updates(&tag).Error
	return
}

func GetTagById(ctx context.Context, id uint) (tag Tag, err error) {
	err = otgorm.WithContext(ctx, db).Select("*").Where(" id = ?", id).First(&tag).Error
	return
}

func GetTagsByIds(ctx context.Context, ids []uint) (tags []Tag, err error) {
	err = otgorm.WithContext(ctx, db).Select("name").Where("id in (?)", ids).Find(&tags).Error
	return
}

func ExitTagWithName(ctx context.Context, name string) (is bool) {
	var tag Tag
	otgorm.WithContext(ctx, db).Select("id").Where(" name = ?", name).First(&tag)
	return tag.ID > 0
}
