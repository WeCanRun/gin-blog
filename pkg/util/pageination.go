package util

import (
	"github.com/WeCanRun/gin-blog/internal/server"
	"github.com/WeCanRun/gin-blog/pkg/setting"
)
import "github.com/unknwon/com"

func GetPage(c *server.Context) (result uint) {
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = uint((page - 1)) * setting.APP.PageSize
	}
	return
}

func GetPageSize(c *server.Context) (res uint) {
	size, _ := com.StrTo(c.Query("page_size")).Int()
	if size <= 0 {
		res = setting.APP.DefaultPageSize
	} else if uint(size) > setting.APP.MaxPageSize {
		res = setting.APP.MaxPageSize
	} else {
		res = uint(size)
	}
	return
}

func GetPageOffset(page, pageSize uint) uint {
	return (page - 1) * pageSize
}
