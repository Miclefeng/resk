package starter

import (
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"github.com/rifflock/lfshook"
	"github.com/lestrrat-go/file-rotatelogs"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/26 上午10:00
 */

func init()  {
	// 日志格式
	formatter := &prefixed.TextFormatter{
		FullTimestamp:   true,
		ForceFormatting: true,
		TimestampFormat: "2006-01-02.15:04:05.000000",
		ForceColors:     true,
		DisableColors:   false,
	}
	logrus.SetFormatter(formatter)
	// 日志级别
	//level := os.Getenv("log.level")
	//if level == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	//}
	// 控制台高亮显示
	formatter.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle: "green",
		DebugLevelStyle: "36",
		WarnLevelStyle: "yellow",
		TimestampStyle: "37",
	})
	// 日志文件和滚动配置
	logrus.Debug("testing")
}