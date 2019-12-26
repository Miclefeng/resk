package starter

import (
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/26 上午10:00
 */

func init()  {
	// 日志格式
	formatter := &logrus.TextFormatter{}
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	logrus.SetFormatter(formatter)
	// 日志级别
	level := os.Getenv("log.level")
	if level == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	// 控制台高亮显示
	formatter.ForceColors = true
	formatter.DisableColors = false
	// 日志文件和滚动配置
}