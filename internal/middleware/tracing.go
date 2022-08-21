package middleware

import (
	"context"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/global/constants"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer, c.Request.URL.Path)
		} else {
			span, ctx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer, c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{
					Key:   string(ext.Component),
					Value: "HTTP",
				})
		}
		defer span.Finish()

		switch span.Context().(type) {
		case jaeger.SpanContext:
			jaegerCtx := span.Context().(jaeger.SpanContext)
			c.Request.Header.Set(constants.TraceId, jaegerCtx.TraceID().String())
			c.Request.Header.Set(constants.SpanId, jaegerCtx.SpanID().String())
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
