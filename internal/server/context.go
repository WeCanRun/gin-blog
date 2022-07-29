package server

import (
	"github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/gin-gonic/gin"
)

type Context struct {
	gin.Context
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
	ctx.JSON(success.StatusCode(), errcode.NewWithData(success.Code(), success.Msg(), data))
	return success
}

func (ctx *Context) ServerError(data interface{}) *errcode.InternalError {
	err := errcode.ServerError
	ctx.AbortWithStatusJSON(err.StatusCode(), errcode.NewWithData(err.Code(), err.Msg(), data))
	return err
}

func (ctx *Context) ParamsError() *errcode.InternalError {
	err := errcode.BadRequest
	ctx.AbortWithStatusJSON(err.StatusCode(), err)
	return err
}
