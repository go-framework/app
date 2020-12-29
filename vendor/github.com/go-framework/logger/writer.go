package logger

import (
	"io"
	"sync"
)

var writerSet sync.Map // map[string]io.Writer

func RegisterWriter(name string, writer io.Writer) {
	writerSet.Store(name, writer)
}

func GetWriter(name string) (io.Writer, bool) {
	value, ok := writerSet.Load(name)
	if !ok {
		return nil, false
	}
	return value.(io.Writer), true
}
