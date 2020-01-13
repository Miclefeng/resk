package starter

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/26 上午10:00
 */

func init() {
	// 日志格式
	stdFormatter := &prefixed.TextFormatter{
		FullTimestamp:   true,
		ForceFormatting: true,
		TimestampFormat: "2006-01-02 15:04:05.000000",
	}

	stdFormatter.ForceColors = true
	stdFormatter.DisableColors = false
	logrus.SetFormatter(stdFormatter)
	// 日志级别
	//level := os.Getenv("log.level")
	//if level == "true" {
	logrus.SetLevel(logrus.DebugLevel)
	//}
	// 控制台高亮显示
	stdFormatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle:  "green",
		DebugLevelStyle: "36",
		WarnLevelStyle:  "yellow",
		TimestampStyle:  "37",
	})
	// 日志文件和滚动配置
	logPath, _ := os.Getwd()
	logNamePre := fmt.Sprintf("%s/log/access_", logPath)

	writer, err := rotatelogs.New(
		logNamePre+"%Y%m%d.log",
		rotatelogs.WithClock(rotatelogs.UTC),
		rotatelogs.WithMaxAge(31*24*time.Hour), // 文件最大保存时间
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
	}, &prefixed.TextFormatter{
		FullTimestamp:   true,
		ForceFormatting: true,
		TimestampFormat: "2006-01-02 15:04:05.000000",
		DisableColors:   true,
	})
	logrus.AddHook(lfHook)
}
