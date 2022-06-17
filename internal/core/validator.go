package core

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
	ja_translations "gopkg.in/go-playground/validator.v9/translations/ja"
	"reflect"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func Init(locale string) (*validator.Validate, error) {
	localeTrans := getLocale(locale)
	uni = ut.New(localeTrans, localeTrans)
	t, _ := uni.GetTranslator(locale)
	trans = t
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		fieldName := fld.Tag.Get("transFieldName")
		if fieldName == "-" {
			return ""
		}
		return fieldName
	})

	if locale == "ja" {
		if e := ja_translations.RegisterDefaultTranslations(validate, trans); e != nil {
			return nil, e
		}
	} else {
		if e := en_translations.RegisterDefaultTranslations(validate, trans); e != nil {
			return nil, e
		}
	}

	return validate, nil
}

func GetErrorMessages(err error) map[string]string {
	if err == nil {
		return map[string]string{}
	}
	return err.(validator.ValidationErrors).Translate(trans)
}

func getLocale(locale string) locales.Translator {
	switch locale {
	case "ja":
		return ja.New()
	default:
		return en.New()
	}
}
