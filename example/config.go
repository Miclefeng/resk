package main

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/23 下午6:00
 */

func main() {
	file := kvs.GetCurrentFilePath("config.ini", 1);
	conf := ini.NewIniFileCompositeConfigSource(file)
	port := conf.GetIntDefault("app.server.port", 18080)
	fmt.Println(file);
	fmt.Println(port);
}
