package starter

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"miclefengzss/resk/bootstrap"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 上午11:15
 */

 var props kvs.ConfigSource

type PropsStarter struct {
	bootstrap.BaseStarter
}

func Props() kvs.ConfigSource {
	return props
}

func (p *PropsStarter) Init(ctx bootstrap.StarterContext) {
	props = ctx.Props()
	fmt.Println("初始化配置。")
}
