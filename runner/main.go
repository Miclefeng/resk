package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "miclefengzss/resk"
	"miclefengzss/resk/bootstrap"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/25 下午4:05
 */

func main() {
	// 获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("../config/config.ini", 1)
	// 加载和解析配置文件
	conf := ini.NewIniFileConfigSource(file)

	app := bootstrap.New(conf)
	app.Start()
}
