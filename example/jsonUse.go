package main

import "github.com/json-iterator/go"
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
	var (
		json jsoniter.ConfigCompatibleWithStandardLibrary
		user User
	)

	user = User{
		Id:      1,
		Name:    "miclefengzss",
		Age:     27,
		Address: "tianan",
	}

	json.Marshal(&user)
}
