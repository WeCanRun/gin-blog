package errcode

import (
	"fmt"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"net/http"
)

var codes = map[int]*InternalError{}

func New(code int, msg string) *InternalError {
	return NewWithData(code, msg, "")
}

func NewWithData(code int, msg string, data interface{}) *InternalError {
	if _, ok := codes[code]; ok {
		logging.Panicf("Code[%d] is exist ", code)
	}

	err := &InternalError{
		Code: code,
		Msg:  msg,
		Data: data,
	}

	codes[code] = err

	return err
}

type InternalError struct {
	Code int         `json:"code" example:"200" example:"400" example:"500" example:"502"`
	Msg  string      `json:"msg" example:"ok"`
	Data interface{} `json:"data,omitempty"`
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

func (e *InternalError) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case BadRequest.Code:
		return http.StatusBadRequest
	case ServerError.Code:
		return http.StatusInternalServerError
	case ErrorAuthCheckTokenFail.Code, ErrorAuthCheckTokenTimeout.Code, ErrorAuthToken.Code, ErrorAuth.Code:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func GetMsg(code int) string {
	e, ok := codes[code]
	if ok {
		return e.Msg
	}
	return ""
}
