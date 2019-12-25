package base

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"miclefengzss/resk/infra"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 上午11:15
 */

 var props kvs.ConfigSource

type PropsStarter struct {
	infra.BaseStarter
}

func Props() (props kvs.ConfigSource) {
	return
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	fmt.Println("初始化配置。")
}
