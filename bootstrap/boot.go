package bootstrap

import "github.com/tietang/props/kvs"

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 上午10:48
 */


// 应用程序启动管理器
type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) (b *BootApplication) {
	b = &BootApplication{conf, StarterContext{}}
	b.starterContext[KeyProps] = conf
	return
}

func (b *BootApplication) Start() {
	// 初始化启动器
	b.init()
	// 安装启动器
	b.setup()
	// 启动启动器
	b.start()
	// 关闭销毁启动器
	b.stop()
}

func (b *BootApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}

func (b *BootApplication) start() {
	for i, starter := range StarterRegister.AllStarters() {
		if starter.StartBlocking() {
			// 如果是最后一个可阻塞的，直接启动阻塞
			if i+1 == len(StarterRegister.starters) {
				starter.Start(b.starterContext)
			} else {
				// 不是最后一个，使用 goroutine 来异步启动后面的启动器
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}

func (b *BootApplication) stop() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Stop(b.starterContext)
	}
}
