package utils

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const phoneRegex = `^01([0|1|6|7|8|9])([0-9]{3,4})([0-9]{4})$`

func ValidateRegex(regex, value string) bool {
	reg := regexp.MustCompile(regex)
	return reg.Match([]byte(value))
}

func RegexPhone(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if value, ok := field.Interface().(string); ok {
		return ValidateRegex(phoneRegex, value)
	}
	return true
}
