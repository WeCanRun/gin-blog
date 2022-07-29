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
