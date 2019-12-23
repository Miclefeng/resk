package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/23 上午10:36
 */

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"UserName"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var (
		user User
		data []byte
	)

	user = User{
		Id:      1,
		Name:    "miclefengzss",
		Age:     27,
		Address: "tianan",
	}

	data, _ = json.Marshal(&user)
	fmt.Println(string(data))
}
