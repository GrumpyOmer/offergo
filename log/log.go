package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

//全局通用log实例
var LogInfo = logrus.New()

// 日志钩子
func NewLfsHook(logName string, rotationTime time.Duration, leastDay uint) logrus.Hook {
	writer, err := rotatelogs.New(
		// 日志文件名称
		logName,

		// 日志周期(默认每86400秒/一天旋转一次)
		rotatelogs.WithRotationTime(rotationTime),

		// 清除历史 (WithMaxAge和WithRotationCount只能选其一)
		//rotatelogs.WithMaxAge(time.Hour*24*7), //默认每7天清除下日志文件
		rotatelogs.WithRotationCount(leastDay), //只保留最近的N个日志文件
	)
	if err != nil {
		panic(err)
	}

	// 可设置按不同level创建不同的文件名
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})

	return lfsHook
}
