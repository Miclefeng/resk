package resk

import (
	"miclefengzss/resk/bootstrap"
	"miclefengzss/resk/bootstrap/starter"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 下午4:08
 */

func init() {
	bootstrap.Register(&starter.PropsStarter{})
	bootstrap.Register(&starter.ValidatorStarter{})
	bootstrap.Register(&starter.DbxDatabaseStarter{})
}
