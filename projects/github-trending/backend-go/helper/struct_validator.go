package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Define a validator struct
type StructValidator struct {
	Validator *validator.Validate
	Uni       *ut.UniversalTranslator
	Trans     ut.Translator
}

// Init an instance of StructValidator struct
func NewStructValidator() *StructValidator {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("en")

	return &StructValidator{
		Validator: validator.New(),
		Uni:       uni,
		Trans:     trans,
	}
}

// Register all validations and translations
func (cv *StructValidator) RegisterValidate() {
	if err := en_translations.RegisterDefaultTranslations(cv.Validator, cv.Trans); err != nil {
		fmt.Errorf(err.Error())
	}

	// Register a custom validation for "pwd" tag. This tag is defined in req_user_signup.go
	cv.Validator.RegisterValidation("pwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) >= 4
	})

	// Register a translation for all required fields (default is english)
	cv.Validator.RegisterTranslation("required", cv.Trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} là bắt buộc", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// Register a translation for "email" field
	cv.Validator.RegisterTranslation("email", cv.Trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} không hợp lệ", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	// Register a translation for "pwd" tag
	cv.Validator.RegisterTranslation("pwd", cv.Trans, func(ut ut.Translator) error {
		return ut.Add("pwd", "Mật khẩu tối thiểu 4 kí tự", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("pwd", fe.Field())
		return t
	})
}

// Override Validate() method so that it can be assigned to echo validator
func (cv *StructValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err == nil {
		return nil
	}

	transErrors := make([]string, 0)
	for _, e := range err.(validator.ValidationErrors) {
		transErrors = append(transErrors, e.Translate(cv.Trans))
	}
	return fmt.Errorf("%s", strings.Join(transErrors, " \n "))
}
