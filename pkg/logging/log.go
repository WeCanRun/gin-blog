package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	File               *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logger             *log.Logger
	logPrefix          = ""
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

func init() {
	filePath := getLogFullPath()
	File, _ = openLogFile(filePath)
	logger = log.New(File, DefaultPrefix, log.LstdFlags)
}

func Debug(fmt string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(fmt+"\n", v)
}

func Info(fmt string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(fmt+"\n", v...)
}

func Warn(fmt string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(fmt+"\n", v...)
}

func Error(fmt string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(fmt+"\n", v...)
}

func Fatal(fmt string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Printf(fmt+"\n", v...)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
