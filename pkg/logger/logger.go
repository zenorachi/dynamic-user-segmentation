package logger

import (
	"log/slog"
	"os"
	"runtime"
	"strconv"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func init() {
	slog.SetDefault(logger)
}

func getCallerInfo() string {
	_, file, line, _ := runtime.Caller(2)
	return file + ":" + strconv.Itoa(line)
}

func Info(service string, args ...interface{}) {
	callerInfo := getCallerInfo()
	slog.Info("unit "+service, append([]interface{}{"caller:" + callerInfo}, args...)...)
}

func Error(service string, args ...interface{}) {
	callerInfo := getCallerInfo()
	slog.Error("service: "+service, append([]interface{}{"caller:" + callerInfo}, args...)...)
}

func Fatal(service string, args ...interface{}) {
	callerInfo := getCallerInfo()
	slog.Error("service: "+service, append([]interface{}{"caller:" + callerInfo}, args...)...)
	os.Exit(1)
}
