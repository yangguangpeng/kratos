package sentry

import (
	"errors"
	"fmt"
)

func Errorf(errorInfoFormat string, a ...interface{}) {
	errorInfo := fmt.Errorf(errorInfoFormat, a...)
	InitSentryInstance.Exception(errorInfo, "Error")
}

func Error(errorInfo error) {
	InitSentryInstance.Exception(errorInfo, "Error")
}

func ErrorString(errorInfo string) {
	InitSentryInstance.Exception(errors.New(errorInfo), "Error")
}

func Exception(errorInfo error) {
	InitSentryInstance.Exception(errorInfo, "Exception")
}

func Info(info interface{}) {
	InitSentryInstance.Info(info, "Info")
}

func Warning(info interface{}) {
	InitSentryInstance.Info(info, "Warning")
}

func Debug(info interface{}) {
	InitSentryInstance.Info(info, "Debug")
}
