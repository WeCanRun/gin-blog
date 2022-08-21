package logging

import (
	"fmt"
	"github.com/WeCanRun/gin-blog/global"
	"github.com/WeCanRun/gin-blog/pkg/file"
	"os"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", global.Setting.APP.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s-%s.%s", global.Setting.APP.LogSaveName, time.Now().Format(global.Setting.APP.LogTimeFormat),
		global.Setting.APP.LogFileExt)
}

func openLogFile(filePath, fileName string) (f *os.File, err error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err:%v", err)
	}
	src := dir + "/" + filePath
	notPerm := file.NotPermission(src)
	if notPerm {
		return nil, fmt.Errorf("file.NotPermission permission denied src: %s", src)
	}
	err = file.IsNotExitMKDir(src)
	if err != nil {
		return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
	}

	f, err = file.Open(src+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("open file fail: %v", err)
	}
	return f, nil
}
