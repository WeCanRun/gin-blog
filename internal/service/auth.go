package service

import (
	e "github.com/WeCanRun/gin-blog/global/errcode"
	"github.com/WeCanRun/gin-blog/internal/model"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/util"
)

func GetTokenWithAuth(username, password string) (code *e.InternalError, token string) {
	ok := model.CheckAuth(username, password)
	if !ok {
		code = e.ErrorAuth
		logging.Error("GetTokenWithAuth#CheckAuth fail")
		return
	}
	var err error
	token, err = util.GenerateToken(username, password)
	if err != nil || len(token) <= 0 {
		code = e.ErrorAuthToken
		logging.Error("GetTokenWithAuth#GenerateToken fail, %v", err)
		return
	}
	return
}
