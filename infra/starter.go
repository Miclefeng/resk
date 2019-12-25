package infra

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
	"reflect"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 上午10:48
 */

const (
	KeyProps = "_conf"
)

// 资源上下文结构体
type StarterContext map[string]interface{}

func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没有被初始化")
	}
	return p.(kvs.ConfigSource)
}

// 基础资源启动接口
type Starter interface {
	// 资源初始化
	Init(ctx StarterContext)
	// 资源的安装
	Setup(ctx StarterContext)
	// 资源启动
	Start(ctx StarterContext)
	// 资源是否需要阻塞启动
	StartBlocking() bool
	// 资源关闭和销毁
	Stop(ctx StarterContext)
}

//基础空启动器实现,为了方便资源启动器的代码实现
var _ Starter = new(BaseStarter)

type BaseStarter struct{}

func (b *BaseStarter) Init(ctx StarterContext)  {}
func (b *BaseStarter) Setup(ctx StarterContext) {}
func (b *BaseStarter) Start(ctx StarterContext) {}
func (b *BaseStarter) StartBlocking() bool      { return false }
func (b *BaseStarter) Stop(ctx StarterContext)  {}

// 服务启动注册器
type starterRegister struct {
	starters []Starter
}

// 启动器注册
func (r *starterRegister) Register(s Starter) {
	r.starters = append(r.starters, s)
	typ := reflect.TypeOf(s)
	logrus.Infof("Register starter: %s", typ.String())
}

// 返回所有启动器
func (r *starterRegister) AllStarters() []Starter {
	return r.starters
}

var StarterRegister *starterRegister = new(starterRegister)

// 外部调用注册函数
func Register(s Starter) {
	StarterRegister.Register(s)
}
