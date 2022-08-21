package global

import (
	"github.com/WeCanRun/gin-blog/pkg/setting"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
)

var (
	DBEngine *gorm.DB
	Tracer   opentracing.Tracer
	Setting  *setting.Setting
)
