package validation

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func NameValidation(fl validator.FieldLevel) bool {
	pattern := `^[a-zA-ZÀ-ÖØ-öø-ÿ'-]+$`
	name := fl.Field().String()
	return regexp.MustCompile(pattern).MatchString(name)
}

func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}
	if !strings.ContainsAny(password, "1234567890") {
		return false
	}
	if !strings.ContainsAny(password, "-!$%@^&*#()_+|~=`{}[]:”;'<>?,./") {
		return false
	}
	return true
}

func NumberPrefixValidation(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return regexp.MustCompile(`^[0-9+]+$`).MatchString(value)
}
