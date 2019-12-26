package starter

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"miclefengzss/resk/bootstrap"
	"time"
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
	bootstrap.BaseStarter
}

func (s *DbxDatabaseStarter) Setup(ctx bootstrap.StarterContext) {
	conf := ctx.Props()
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	settings.ConnMaxLifetime = conf.GetDurationDefault("mysql.connMaxLifetime", 8*time.Hour)
	settings.Options = map[string]string{
		"charset": conf.GetDefault("mysql.charset", "utf8mb4"),
		"parseTime": conf.GetDefault("mysql.parseTime", "true")}
	logrus.Infof("mysql.con url: %s", settings.ShortDataSourceName())
	database, err = dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info("mysql.ping: ", database.Ping())
}
