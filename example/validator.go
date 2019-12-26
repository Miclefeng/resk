package main

import (
	"fmt"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/locales/zh"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/validator.v9/translations/zh"
)

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

func main() {
	user := &Users{
		FirstName: "zhang",
		LastName:  "ShangShang",
		Age:       -1,
		Email:     "miclefengzss163.com",
	}

	validate := validator.New()

	cn := zh.New()
	uni := ut.New(cn, cn)

	err := validate.Struct(user)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			fmt.Println(err)
		}
		errors, ok := err.(validator.ValidationErrors)
		if ok {
			fmt.Println(errors)
		}
	}
}
