package resk

import (
	"resk/infra"
	"resk/infra/base"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 下午4:08
 */

func init() {
	infra.Register(&base.PropsStarter{})
}
