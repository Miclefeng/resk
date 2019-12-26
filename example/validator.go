package main

import "gopkg.in/go-playground/validator.v9"
/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/26 下午3:46
 */

type Users struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       int    `validate:"gte=0,lte=150"`
	Email     string `validate:"required,email"`
}

func main()  {
	user := &Users{
		FirstName: "zhang",
		LastName: "ShangShang",
		Age: 27,
		Email: "miclefengzss@163.com",
	}

	validate := validator.New()
}