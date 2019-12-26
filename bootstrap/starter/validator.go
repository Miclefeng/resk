package starter

import (
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/locales/zh"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/go-playground/validator.v9/translations/zh"
	"miclefengzss/resk/bootstrap"
)

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/12/26 下午5:05
 */

var (
	validate   *validator.Validate
	translator ut.Translator
)

func Validate() *validator.Validate {
	return validate
}

func Translate() ut.Translator {
	return translator
}

type ValidateStarter struct {
	bootstrap.BaseStarter
}

func (v *ValidateStarter) Init(ctx bootstrap.Starter) {
	validate = validator.New()
	cn := zh.New()
	translator := ut.New(cn, cn)

}