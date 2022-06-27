package utils

import (
	"log"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func InitLog() {
	//日志文件
	logDir := "logs"
	os.Mkdir(logDir, os.ModePerm)
	fileName := path.Join(logDir, "agent.log")

	//写入文件
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		log.Fatalln("日志文件访问有错误：", err)
	}

	Logger = logrus.New()
	Logger.Out = file
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.JSONFormatter{})

	//var maxRemainCnt uint = 10
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d%H.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(每小时)
		rotatelogs.WithMaxAge(time.Hour),
		// WithMaxAge和WithRotationCount二者只能设置一个
		// WithMaxAge设置文件清理前的最长保存时间
		// WithRotationCount设置文件清理前最多保存的个数
		rotatelogs.WithMaxAge(time.Hour*24*7),
		//rotatelogs.WithRotationCount(maxRemainCnt),
		// 设置日志切割时间间隔(1天)
		//rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

}
