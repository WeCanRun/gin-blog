package e

import (
	"github.com/WeCanRun/gin-blog/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	SERVER_ERROR:                   "系统错误",
	ERROR_EXIST_TAG:                "已存在该标签名称",
	ERROR_NOT_EXIST_TAG:            "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

func OtherError(ctx *gin.Context, code int)  {
	ctx.AbortWithStatusJSON(http.StatusBadGateway, dto.Response{
		Code: code,
		Msg:  GetMsg(code),
		Data: nil,
	})
}

func AuthError(ctx *gin.Context, code int)  {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
		Code: code,
		Msg:  GetMsg(code),
		Data: nil,
	})
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(SUCCESS, dto.Response{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: data,
	})
}

func ServerError(ctx *gin.Context, data interface{}) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
		Code: SERVER_ERROR,
		Msg:  GetMsg(SERVER_ERROR),
		Data: data,
	})
}

func ParamsError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
		Code: INVALID_PARAMS,
		Msg:  GetMsg(INVALID_PARAMS),
		Data: nil,
	})
}
