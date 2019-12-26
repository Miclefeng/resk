package starter

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
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

type ValidatorStarter struct {
	bootstrap.BaseStarter
}

func (v *ValidatorStarter) Init(ctx bootstrap.StarterContext) {
	validate = validator.New()
	cn := zh.New()
	uni := ut.New(cn, cn)
	var found bool
	translator, found = uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			logrus.Error(err)
		}
	} else {
		logrus.Error("Not found: translator zh")
	}
}
