package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	File               *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 4
	l                  *Logger
	logPrefix          = ""
)

type Level int

const (
	DEBUG Level = iota + 1
	INFO
	WARNING
	ERROR
	FATAL
	PANIC
)

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return ""
	}
}

func Setup() {
	fileName := getLogFilePath() + getLogFileName()
	output := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    600,
		MaxAge:     10,
		MaxBackups: 0,
		LocalTime:  true,
	}
	writer := io.MultiWriter(output, os.Stdout)
	l = New(writer, DefaultPrefix, log.LstdFlags)
}

func Log() *Logger {
	return l
}

type Fields map[string]interface{}

type Logger struct {
	logger  *log.Logger
	ctx     context.Context
	fields  Fields
	callers []string
}

func New(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{logger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	caller, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(caller)
		split := strings.Split(f.Name(), ".")
		name := split[len(split)-1]
		ll.callers = []string{fmt.Sprintf("%s:%d (%s)", file, line, name)}
	}
	return ll
}

func (l *Logger) WithCallerFrames() *Logger {
	var callers []string
	minDepth, maxDepth := 1, 30
	pcs := make([]uintptr, maxDepth)
	depth := runtime.Callers(minDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for next, more := frames.Next(); more; next, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", next.File, next.Line, next.Function)
		callers = append(callers, s)
	}

	ll := l.clone()
	ll.callers = callers
	return ll
}

func (l *Logger) JSONFormat(level Level, message string) Fields {
	data := make(Fields, len(l.fields)+3)
	data["level"] = level.String()
	data["content"] = message
	l.callers = l.WithCaller(DefaultCallerDepth).callers
	data["callers"] = l.callers

	for k, v := range l.fields {
		if _, ok := data[k]; !ok {
			data[k] = v
		}
	}
	return data
}

func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case DEBUG, INFO, WARNING, ERROR:
		l.logger.Println(content)
	case FATAL:
		l.logger.Fatalln(content)
	case PANIC:
		l.logger.Panicln(content)
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.Output(DEBUG, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(DEBUG, fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.Output(INFO, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(INFO, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.Output(WARNING, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(WARNING, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(ERROR, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(ERROR, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(FATAL, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(FATAL, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.Output(PANIC, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Output(PANIC, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	l.Output(DEBUG, fmt.Sprint(v...))
}
func Debugf(format string, v ...interface{}) {
	l.Output(DEBUG, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	l.Output(INFO, fmt.Sprint(v...))
}
func Infof(format string, v ...interface{}) {
	l.Output(INFO, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	l.Output(WARNING, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	l.Output(WARNING, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	l.Output(ERROR, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	l.Output(ERROR, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	l.Output(FATAL, fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) {
	l.Output(FATAL, fmt.Sprintf(format, v...))
}

func Panic(v ...interface{}) {
	l.Output(PANIC, fmt.Sprint(v...))

}

func Panicf(format string, v ...interface{}) {
	l.Output(WARNING, fmt.Sprintf(format, v...))
}
