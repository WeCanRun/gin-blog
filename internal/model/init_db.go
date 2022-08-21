package model

import (
	"fmt"
	otgorm "github.com/EDDYCJY/opentracing-gorm"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Setup() {
	var err error
	dbStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.Database.User,
		setting.Database.Password,
		setting.Database.Host,
		setting.Database.DbName)
	db, err = gorm.Open(setting.Database.Type, dbStr)
	if err != nil || db == nil {
		logging.Log().Fatal("db init fail, dbStr: %s, err: %v", dbStr, err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.Database.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	AutoBuildTable()
	otgorm.AddGormCallbacks(db)

}

func AutoBuildTable() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Article{}, &Tag{}, &Auth{})
}

func CloseDB() {
	defer db.Close()
}
