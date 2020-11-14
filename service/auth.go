package service

import (
	"github.com/WeCanRun/gin-blog/model"
	"github.com/WeCanRun/gin-blog/pkg/e"
	"github.com/WeCanRun/gin-blog/pkg/logging"
	"github.com/WeCanRun/gin-blog/pkg/util"
)

func GetTokenWithAuth(username, password string) (code int, token string) {
	ok := model.CheckAuth(username, password)
	if !ok {
		code = e.ERROR_AUTH
		logging.Error("GetTokenWithAuth#CheckAuth fail")
		return
	}
	var err error
	token, err = util.GenerateToken(username, password)
	if err != nil || len(token) <= 0 {
		code = e.ERROR_AUTH_TOKEN
		logging.Error("GetTokenWithAuth#GenerateToken fail, %v", err)
		return
	}
	return
}
