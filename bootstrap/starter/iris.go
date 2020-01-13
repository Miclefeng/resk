package starter

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisRecover "github.com/kataras/iris/middleware/recover"
	"github.com/sirupsen/logrus"
	"miclefengzss/resk/bootstrap"
	"time"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/30 上午11:21
 */

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisStarter struct {
	bootstrap.BaseStarter
}

func (i *IrisStarter) Init(ctx bootstrap.StarterContext) {
	// 创建irisApplication实例
	irisApplication = initIris()
	// 日志组件配置和扩展
	logs := irisApplication.Logger()
	logs.Install(logrus.StandardLogger())
}

func (i *IrisStarter) Start(ctx bootstrap.StarterContext) {
	// 把路由信息打印到控制台
	routes := irisApplication.GetRoutes()
	for _, route := range routes {
		logrus.Info(route.Trace())
	}
	// 启动iris
	port := ctx.Props().GetDefault("app.server.port", "18080")
	irisApplication.Run(iris.Addr(":" + port))
}

func (i *IrisStarter) StartBlocking() bool {
	return true
}

func initIris() *iris.Application {
	app := iris.New()

	// 主要中间件配置：recover 和 日志输出中间件自定义
	app.Use(irisRecover.New())
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		LogFunc: func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
			app.Logger().Infof("%s | %s | %s | %s | %s | %s | %s | %s", now.Format("2006-01-02 15:04:05.000"), latency.String(), status, ip, method, path, message, headerMessage)
		},
	}
	app.Use(logger.New(cfg))
	return app
}
