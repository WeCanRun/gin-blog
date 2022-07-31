package server

import (
	"github.com/WeCanRun/gin-blog/global/constants"
	"github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type Context struct {
	*gin.Context
	traceId string
	spanId  string
	logger  *logging.Logger
}

func (ctx *Context) OtherError(code int) *errcode.InternalError {
	err := errcode.New(code, errcode.GetMsg(code))
	ctx.AbortWithStatusJSON(err.StatusCode(), err)
	return err
}

func (ctx *Context) AuthError() *errcode.InternalError {
	err := errcode.ErrorAuth
	ctx.AbortWithStatusJSON(err.StatusCode(), err)
	return err
}

func (ctx *Context) Success(data interface{}) *errcode.InternalError {
	success := errcode.Success
	ctx.JSON(success.StatusCode(), errcode.NewWithData(success.Code, success.Msg, data))
	return success
}

func (ctx *Context) SuccessList(data interface{}) *errcode.InternalError {
	success := errcode.Success
	ctx.JSON(success.StatusCode(), errcode.NewWithData(success.Code, success.Msg, data))
	return success
}

func (ctx *Context) ServerError(data interface{}) *errcode.InternalError {
	err := errcode.ServerError
	ctx.AbortWithStatusJSON(err.StatusCode(), errcode.NewWithData(err.Code, err.Msg, data))
	return err
}

func (ctx *Context) ParamsError() *errcode.InternalError {
	err := errcode.BadRequest
	ctx.AbortWithStatusJSON(err.StatusCode(), err)
	return err
}

func (ctx *Context) Logger() *logging.Logger {
	return ctx.logger
}

type Handler func(*Context) error

func HandlerWarp(handler ...Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		traceId := ctx.Request.Header.Get(constants.TraceId)
		if len(traceId) == 0 {
			traceId = uuid.New().String()
			ctx.Request.Header.Set(constants.TraceId, traceId)
		}

		spanId := ctx.Request.Header.Get(constants.SpanId)
		if len(spanId) == 0 {
			spanId = uuid.New().String()
			ctx.Request.Header.Set(constants.SpanId, spanId)
		}

		customCtx := &Context{
			Context: ctx,
			traceId: traceId,
			spanId:  spanId,
			logger:  logging.Log().WithFields(map[string]interface{}{"traceId": traceId, "spanId": spanId}),
		}

		var err error
		for _, h := range handler {
			if err = h(customCtx); err != nil {
				_, ok := err.(*errcode.InternalError)
				if !ok {
					err = customCtx.ServerError(err.Error())
				}
			}
		}

		spend := time.Now().Sub(start).Milliseconds()
		logging.Infof("Spend time: %d, response: %#v", spend, err)
	}
}
