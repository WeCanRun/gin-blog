package server

import (
	"github.com/WeCanRun/gin-blog/global/constants"
	"github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Context struct {
	*gin.Context
	traceId string
	spanId  string
	logger  *logging.Logger
}

func (ctx *Context) WithContext(c *gin.Context) *Context {
	if c == nil {
		panic("nil context")
	}
	ctx2 := new(Context)
	*ctx2 = *ctx
	ctx2.Context = c
	return ctx2
}

func (ctx *Context) OtherError(code int) *errcode.InternalError {
	err := errcode.New(code, errcode.GetMsg(code))
	return err
}

func (ctx *Context) AuthError() *errcode.InternalError {
	err := errcode.ErrorAuth
	return err
}

func (ctx *Context) Success(data interface{}) *errcode.InternalError {
	success := errcode.Success
	success.Data = data
	return success
}

func (ctx *Context) SuccessList(data interface{}) *errcode.InternalError {
	success := errcode.Success
	success.Data = data
	return success
}

func (ctx *Context) ServerError(data interface{}) *errcode.InternalError {
	err := errcode.ServerError
	err.Data = data
	return err
}

func (ctx *Context) ParamsError(data interface{}) *errcode.InternalError {
	err := errcode.BadRequest
	err.Data = data
	return err
}

func (ctx *Context) Logger() *logging.Logger {
	return ctx.logger
}

func (ctx *Context) Bind(v interface{}) errcode.ValiadErrors {
	var res errcode.ValiadErrors
	if err := ctx.Context.Bind(v); err != nil {
		value := ctx.Value(constants.Trans)
		trans, _ := value.(ut.Translator)
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return res
		}
		for k, v := range errs.Translate(trans) {
			res = append(res, &errcode.ValiadError{
				Key: k,
				Msg: util.CamelCaseToUnderscore(v),
			})
		}

		return res
	}
	return nil
}

type Handler func(*Context) error

func HandlerWarp(handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceId := ctx.Request.Header.Get(constants.TraceId)
		if len(traceId) == 0 {
			traceId = uuid.New().String()[:8]
			ctx.Request.Header.Set(constants.TraceId, traceId)
		}

		spanId := ctx.Request.Header.Get(constants.SpanId)
		if len(spanId) == 0 {
			spanId = uuid.New().String()[:8]
			ctx.Request.Header.Set(constants.SpanId, spanId)
		}

		customCtx := &Context{
			Context: ctx,
			traceId: traceId,
			spanId:  spanId,
			logger: logging.Log().WithFields(map[string]interface{}{
				constants.LogFieldTraceId: ctx.Request.Header.Get(constants.TraceId),
				constants.LogFieldSpanId:  ctx.Request.Header.Get(constants.SpanId)}),
		}

		customCtx.Writer.Header().Set("content-type", "text/json; charset=utf-8")
		if err := handler(customCtx); err == nil {
			customCtx.JSON(errcode.Success.StatusCode(), errcode.Success)

		} else {
			ierr, ok := err.(*errcode.InternalError)
			if !ok {
				ierr = customCtx.ServerError(err.Error())
			}

			if ierr != nil {
				ctx.JSON(ierr.StatusCode(), ierr)
			} else {
				customCtx.Logger().Warn("Should not return nil")
			}
		}

		customCtx.Logger().Debug("HandlerWarp")
	}
}
