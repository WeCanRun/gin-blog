package global

import (
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
)

var (
	DBEngine *gorm.DB
	Tracer   opentracing.Tracer
)
