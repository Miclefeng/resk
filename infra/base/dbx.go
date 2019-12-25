package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"miclefengzss/resk/infra"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 下午5:11
 */

var database *dbx.Database

func DbxDatabase() *dbx.Database {
	return database
}

type DbxDatabaseStarter struct {
	infra.BaseStarter
}

func (s *DbxDatabaseStarter) Setup(ctx infra.StarterContext) {
	conf := ctx.Props()
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}

	logrus.Debug("mysql.con url: ", settings.ShortDataSourceName())
}
