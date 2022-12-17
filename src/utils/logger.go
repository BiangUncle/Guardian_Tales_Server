package utils

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

const LogDirName = "logs/"

// Init 初始化日志框架
func Init() {
	// 获取工作路径
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 设置颜色
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	// 设置文件名
	timeFormat := time.Now().Format("2006-01-02 15:04:05")
	// 设置输出路径
	file := dir + "/" + LogDirName + timeFormat + ".log"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	// 输出到文件和STDOUT
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(fileAndStdoutWriter) // 将文件设置为log输出的文件
}
