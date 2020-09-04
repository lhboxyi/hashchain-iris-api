package config

import (
	"bytes"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"iris-api/com/util"
	"os"
	"strings"
	"time"
)

/**
日志
*/
func NewLfsHook() logrus.Hook {
	cnt := viper.Get("logsetting.maxRemainCnt").(int)
	path := viper.Get("logsetting.fileDir").(string)
	logName := viper.Get("logsetting.fileName").(string)
	logLevel := viper.Get("logsetting.logLevel").(string)
	//如果目录不存在，则创建该目录
	if !util.Exists(path) {
		os.MkdirAll(path, os.ModePerm)
	}
	writer, err := rotatelogs.New(
		path+"/"+logName+"%Y%m%d%H"+".log",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		//rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(24*time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(uint(cnt)),
	)

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	//logrus.SetLevel(logrus.WarnLevel)
	level, _ := logrus.ParseLevel(strings.ToLower(logLevel))
	logrus.SetLevel(level)
	logrus.SetFormatter(new(Formatter))

	lfsHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: writer,
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		}, //&logrus.TextFormatter{DisableColors: true}
		&logrus.JSONFormatter{})

	return lfsHook
}

type Formatter struct{}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var out string
	if entry.HasCaller() {
		var pkg bytes.Buffer
		for _, t := range strings.Split(entry.Caller.File, "/") {
			if len(t) == 0 {
				continue
			}
			pkg.WriteString(t[:1])
			pkg.WriteString(".")
		}
		pkg.WriteString(entry.Caller.Function)

		out = fmt.Sprintf("[%s][%s] %s(%d) %s\n", entry.Level.String()[:4], entry.Time.Format("2006-01-02 15:04:05"), pkg.String(), entry.Caller.Line, entry.Message)
	} else {
		out = fmt.Sprintf("[%s][%s] %s\n", entry.Level.String()[:4], entry.Time.Format("2006-01-02 15:04:05"), entry.Message)
	}
	return []byte(out), nil
}
