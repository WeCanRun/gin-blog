package middleware

import (
	"bytes"
	"github.com/WeCanRun/gin-blog/global/constants"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}

	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		logging.Debug("Use AccessLog")
		bw := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}

		c.Writer = bw

		begin := time.Now().Unix()
		c.Next()
		end := time.Now().Unix()

		fields := logging.Fields{
			constants.LogFieldTraceId: c.Request.Header.Get(constants.TraceId),
			constants.LogFieldSpanId:  c.Request.Header.Get(constants.SpanId),
			"request":                 c.Request.PostForm.Encode(),
			"response":                bw.body.String(),
		}

		s := "access log: method: %s, status_code: %d, begin_time:%d, end_time: %d"

		logging.Log().WithFields(fields).Infof(s, c.Request.Method, bw.Status(), begin, end)
	}
}
