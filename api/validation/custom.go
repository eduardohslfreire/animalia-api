package validation

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Translator ...
var Translator ut.Translator

func init() {
	en := en.New()
	Universal := ut.New(en, en)

	Translator, _ = Universal.GetTranslator("en")
}

// RegisterCustomValidations ...
func RegisterCustomValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTranslation("required", Translator, func(ut ut.Translator) error {
			return ut.Add("required", "{0} must have a value", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())

			return t
		})
	}
}
