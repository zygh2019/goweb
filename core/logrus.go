package core

import (
	"awesomeProject1/globle"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 34
	gray   = 37
)

type LogFormatter struct{}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {

		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	loggerConfig := globle.Config.Logger
	if entry.HasCaller() {
		//自定义文件路径
		funcval := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", loggerConfig.Prefix, timestamp, levelColor, entry.Level, fileVal, funcval, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s][%s] \x1b[%dm[%s]\x1b[0m %s\n", loggerConfig.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mlog := logrus.New()               //新建一个实例
	mlog.SetOutput(os.Stdout)          //设置输出类型
	mlog.SetFormatter(&LogFormatter{}) //设置自定义哥格式化
	level, error := logrus.ParseLevel(globle.Config.Logger.Level)
	if error != nil {
		level = logrus.InfoLevel
	}
	mlog.SetLevel(level)                                //设置最低的level界别
	mlog.SetReportCaller(globle.Config.Logger.ShowLine) //开启返回函数名和行号
	initDefaultLogger()
	globle.Log = mlog
	return mlog
}

func initDefaultLogger() {

	logrus.SetFormatter(&LogFormatter{})
	logrus.SetOutput(os.Stdout)          //设置输出类型
	logrus.SetFormatter(&LogFormatter{}) //设置自定义哥格式化
	level, error := logrus.ParseLevel(globle.Config.Logger.Level)
	if error != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true) //开启返回函数名和行号

}
