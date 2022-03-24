package rest

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	Validator *validator.Validate
	Translator ut.Translator
}


func NewValidator() (*Validator, error) {
	uni := ut.New(en.New())
	translation, _ := uni.GetTranslator("en")
	validate := validator.New()

	err := en_translations.RegisterDefaultTranslations(validate, translation)
	if err != nil {
		return nil, err
	}
	return &Validator{validate, translation} , nil
}
