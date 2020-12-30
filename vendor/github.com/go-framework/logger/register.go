package logger

import (
	"os"

	"github.com/natefinch/lumberjack"

	file_rotatelogs "github.com/go-framework/logger/writers/file-rotatelogs"
)

func init() {
	RegisterWriter("console", os.Stdout)
	RegisterWriter("lumberjack", &lumberjack.Logger{})
	RegisterWriter("file-rotatelogs", &file_rotatelogs.FileRotateLogs{})
}
