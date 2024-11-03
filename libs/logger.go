package libs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
	"time"
)

var hasInit = 0
var logger = logrus.New()
var config, _ = LoadConfig()

type NoFieldNamesFormatter struct {
	TimestampFormat string
}

func (f *NoFieldNamesFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	parts := strings.Split(entry.Caller.File, "/")
	filename := parts[len(parts)-1]
	formatted := fmt.Sprintf("[%s] %-5s | %s:%d ==> %s\n",
		entry.Time.Format("2006-01-02 15:04:05.000"),
		strings.ToUpper(entry.Level.String()),
		filename,
		entry.Caller.Line,
		entry.Message)

	return []byte(formatted), nil
}

func InitLogger() {
	logger.SetReportCaller(true)
	logger.SetFormatter(&NoFieldNamesFormatter{})

	os.MkdirAll(config.LogPath, os.ModePerm)

	file := lumberjack.Logger{
		Filename:   config.LogPath + "/viz.log",
		MaxSize:    10, // megabytes
		MaxBackups: 7,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	}
	go func() {
		for {
			nowTime := time.Now()
			nowTimeStr := nowTime.Format("2006-01-02")
			t2, _ := time.ParseInLocation("2006-01-02", nowTimeStr, time.Local)
			next := t2.AddDate(0, 0, 1)
			after := next.UnixNano() - nowTime.UnixNano() - 1
			<-time.After(time.Duration(after) * time.Nanosecond)
			file.Rotate()
		}
	}()

	logger.SetOutput(io.MultiWriter(&file, os.Stdout))

}

func GetLogger() *logrus.Logger {
	if hasInit == 0 {
		InitLogger()
	}
	hasInit = 1
	return logger
}
