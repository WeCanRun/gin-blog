package model

import (
	"github.com/jinzhu/gorm"
	"time"
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
func GetTags(pageNum uint, pageSize uint) (tags []Tag, err error) {
	err = db.Where("state = ?", 1).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	return
}

func GetTagsByName(name string) (tags []Tag, err error) {
	err = db.Model(Tag{}).Where("name = ?", name).Find(&tags).Error
	return
}

func DeleteTag(id uint) (err error) {
	return db.Model(Tag{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func AddTag(tag Tag) (err error) {
	err = db.Model(tag).Create(&tag).Error
	return
}

func EditTag(tag Tag) (err error) {
	err = db.Model(tag).Where("id = ?", tag.ID).Updates(&tag).Error
	return
}

func GetTagById(id uint) (tag Tag, err error) {
	err = db.Select("*").Where(" id = ?", id).First(&tag).Error
	return
}

func GetTagsByIds(ids []uint) (tags []Tag, err error) {
	err = db.Select("name").Where("id in (?)", ids).Find(&tags).Error
	return
}

func ExitTagWithName(name string) (is bool) {
	var tag Tag
	db.Select("id").Where(" name = ?", name).First(&tag)
	return tag.ID > 0
}
