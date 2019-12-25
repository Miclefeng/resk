package brun

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "miclefengzss/resk"
	"miclefengzss/resk/infra"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 下午4:05
 */

func main() {
	// 获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("config.ini", 1)
	// 加载和解析配置文件
	conf := ini.NewIniFileConfigSource(file)

	app := infra.New(conf)
	app.Start()
}
