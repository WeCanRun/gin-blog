package middleware

import (
	"github.com/WeCanRun/gin-blog/global/constants"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	en_validator "github.com/go-playground/validator/v10/translations/en"
	zh_validator "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locate")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "en":
				_ = en_validator.RegisterDefaultTranslations(v, trans)
			case "zh":
				fallthrough
			default:
				_ = zh_validator.RegisterDefaultTranslations(v, trans)
			}

			c.Set(constants.Trans, trans)
		}
		c.Next()
	}
}
