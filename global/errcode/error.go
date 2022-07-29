package errcode

import (
	"fmt"
	"net/http"
)

var codes = map[int]*InternalError{}

func New(code int, msg string) *InternalError {
	return NewWithData(code, msg, "")
}

func NewWithData(code int, msg string, data interface{}) *InternalError {
	if e, ok := codes[code]; ok {
		e.msg = msg
		e.data = data
		return e
	}

	err := &InternalError{
		code: code,
		msg:  msg,
		data: data,
	}

	return err
}

type InternalError struct {
	code int         `json:"code" example:"200" example:"400" example:"500" example:"502"`
	msg  string      `json:"msg" example:"ok"`
	data interface{} `json:"data,omitempty"`
}

func (e *InternalError) Error() string {
	return fmt.Sprintf("%d: %s", e.code, e.msg)
}

func (e *InternalError) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case BadRequest.code:
		return http.StatusBadRequest
	case ServerError.code:
		return http.StatusInternalServerError
	case ErrorAuthCheckTokenFail.code, ErrorAuthCheckTokenTimeout.code, ErrorAuthToken.code, ErrorAuth.code:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func (e *InternalError) Code() int {
	return e.code
}

func (e *InternalError) Msg() string {
	return e.msg
}

func (e *InternalError) Data() interface{} {
	return e.data
}

func GetMsg(code int) string {
	e, ok := codes[code]
	if ok {
		return e.msg
	}
	return ""
}
