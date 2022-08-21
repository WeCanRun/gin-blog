package model

import (
	"fmt"
	otgorm "github.com/EDDYCJY/opentracing-gorm"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func Setup() {
	var err error
	dbStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.Setting.Database.User,
		global.Setting.Database.Password,
		global.Setting.Database.Host,
		global.Setting.Database.DbName)
	db, err = gorm.Open(global.Setting.Database.Type, dbStr)
	if err != nil || db == nil {
		logging.Log().Fatalf("db init fail, dbStr: %s, err: %v", dbStr, err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return global.Setting.Database.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(global.Setting.Database.MaxIdleConns)
	db.DB().SetMaxOpenConns(global.Setting.Database.MaxOpenConns)

	AutoBuildTable()
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", DeleteCallback)
	otgorm.AddGormCallbacks(db)
}

func DeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var str string
		if field, ok := scope.Get("gorm:delete_option"); ok {
			str = field.(string)
		}
		deleteAt, hasDeleteAt := scope.FieldByName("DeletedAt")
		if !scope.Search.Unscoped && hasDeleteAt {
			now := time.Now()
			scope.Raw(fmt.Sprintf(
				"update %v set %v=%v %v %v",
				scope.QuotedTableName(),
				scope.Quote(deleteAt.DBName),
				scope.AddToVars(now),
				scope.CombinedConditionSql(),
				str)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"delete from %v %v %v",
				scope.QuotedTableName(),
				scope.CombinedConditionSql(),
				str)).Exec()
		}

	}
}

func updateTimeForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now()
		if _, ok := scope.Get("gorm:update_column"); !ok {
			_ = scope.SetColumn("UpdatedAt", now)
		}
	}
}

func updateTimeForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		now := time.Now()
		if create, ok := scope.FieldByName("CreatedAt"); ok {
			if create.IsBlank {
				_ = create.Set(now)
			}
		}

		if update, ok := scope.FieldByName("UpdatedAt"); ok {
			if update.IsBlank {
				_ = update.Set(now)
			}
		}
	}
}

func AutoBuildTable() {
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Article{}, &Tag{}, &Auth{})
}

func CloseDB() {
	defer db.Close()
}
