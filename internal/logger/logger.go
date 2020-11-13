package logger

import (
	"github.com/fatih/color"
	"os"
)

func Error(format string, v ...interface{}) {
	color.Red(format, v...)
	os.Exit(1)
}

func Info(format string, v ...interface{}) {
	color.Cyan(format, v...)
}
