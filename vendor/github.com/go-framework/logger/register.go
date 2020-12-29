package logger

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
)

func init() {
	RegisterWriter("console", os.Stdout)
	RegisterWriter("lumberjack", &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s.log", os.Args[0]),
		MaxSize:    512,
		MaxBackups: 5,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	})
}
