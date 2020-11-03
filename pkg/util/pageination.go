package util

import (
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)
import "github.com/unknwon/com"

func GetPage(c *gin.Context) (result uint) {
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = uint((page - 1)) * setting.App.PageSize
	}
	return
}
