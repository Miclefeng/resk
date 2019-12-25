package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/23 下午6:00
 */

func main() {
	file := kvs.GetCurrentFilePath("config.ini", 1);
	fmt.Printfln(file);
}
