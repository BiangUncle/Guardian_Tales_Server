package logger

import (
	"bytes"
	"log"
	"sync"
)

var _logger *log.Logger
var _loggerOnce sync.Once
var buf bytes.Buffer

func Init() {
	//_loggerOnce.Do(func() {
	//	_logger = log.New(&buf, "", log.Lshortfile)
	//})
	//return _logger
	_logger = log.New(&buf, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}

func Info(msg string) {
	buf.Reset()
	_logger.SetPrefix("[INFO]: ")
	_logger.Output(1, msg)
}

func Warn(msg string) {
	buf.Reset()
	_logger.SetPrefix("[Warn]: ")
	_logger.Output(1, msg)
}

func Error(msg string) {
	buf.Reset()
	_logger.SetPrefix("[Error]: ")
	_logger.Output(1, msg)
}
