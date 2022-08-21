package middleware

import (
	"bytes"
	"fmt"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/global/constants"
	e "github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/pkg/file"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logging.Error(err)
				res := e.ServerError
				res.Data = err

				subject := fmt.Sprintf("抛出异常： 发送时间: %v", time.Now().Format("2006-01-02 15:04:05	"))
				fields := logging.Fields{
					"Subject":  subject,
					"TraceId":  c.Request.Header.Get(constants.TraceId),
					"SpanId":   c.Request.Header.Get(constants.SpanId),
					"Request":  fmt.Sprintf("%#v", c.Request.PostForm.Encode()),
					"Response": fmt.Sprintf("%#v", res),
				}

				var body bytes.Buffer
				name := file.CoverToAbs("conf/email.html")
				t, err := template.ParseFiles(name)
				if err != nil {
					logging.Errorf("parse %s fail, err: %v", name, err)
				}

				err = t.Execute(&body, fields)
				if err != nil {
					logging.Errorf("execute fail, err: %v", err)
				}

				util.SendEmail(global.Setting.Email.To,
					subject,
					body.String())

				c.AbortWithStatusJSON(res.StatusCode(), res)
			}
		}()
		c.Next()
	}
}
