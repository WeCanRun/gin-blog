package util

import (
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/internal/server"
)
import "github.com/unknwon/com"

func GetPage(c *server.Context) (result uint) {
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = uint((page - 1)) * global.Setting.APP.PageSize
	}
	return
}

func GetPageSize(c *server.Context) (res uint) {
	size, _ := com.StrTo(c.Query("page_size")).Int()
	if size <= 0 {
		res = global.Setting.APP.DefaultPageSize
	} else if uint(size) > global.Setting.APP.MaxPageSize {
		res = global.Setting.APP.MaxPageSize
	} else {
		res = uint(size)
	}
	return
}

func GetPageOffset(page, pageSize uint) uint {
	return (page - 1) * pageSize
}
