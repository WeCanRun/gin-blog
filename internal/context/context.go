package context

import (
	"github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	gin.Context
}

func (ctx *Context) OtherError(code int) {
	ctx.AbortWithStatusJSON(http.StatusBadGateway, dto.Response{
		Code: code,
		Msg:  errcode.GetMsg(code),
		Data: nil,
	})
}

func (ctx *Context) AuthError(code int) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
		Code: code,
		Msg:  errcode.GetMsg(code),
		Data: nil,
	})
}

func (ctx *Context) Success(data interface{}) {
	ctx.JSON(errcode.SUCCESS, dto.Response{
		Code: errcode.SUCCESS,
		Msg:  errcode.GetMsg(errcode.SUCCESS),
		Data: data,
	})
}

func (ctx *Context) ServerError(data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
		Code: errcode.SERVER_ERROR,
		Msg:  errcode.GetMsg(errcode.SERVER_ERROR),
		Data: data,
	})
}

func (ctx *Context) ParamsError() {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
		Code: errcode.INVALID_PARAMS,
		Msg:  errcode.GetMsg(errcode.INVALID_PARAMS),
		Data: nil,
	})
}
