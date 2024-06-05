package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func ValidateErrors[T any](err error, arg T) (verr []string) {
	if _, ok := err.(validator.ValidationErrors); ok {
		for _, err := range err.(validator.ValidationErrors) {
			var e string
			fieldName := getJsonTag(err.Field(), arg)
			switch err.Tag() {
			case "required":
				e = fmt.Sprintf("field '%s' is required", fieldName)
			case "lte":
				e = fmt.Sprintf("field '%s' must be less or equal '%v' instead of '%v'", fieldName, err.Param(), err.Value())
			case "gt":
				e = fmt.Sprintf("field '%s' must be greather than '%v'", fieldName, err.Param())
			case "min":
				e = fmt.Sprintf("field '%s' must have at least '%v' characters", fieldName, err.Param())
			case "max":
				e = fmt.Sprintf("field '%s' must have at most '%v' characters", fieldName, err.Param())
			default:
				e = fmt.Sprintf("field '%s': '%v' must satisfy '%s' '%v' criteria", fieldName, err.Value(), err.Tag(), err.Param())
			}
			verr = append(verr, e)
		}
	}
	return verr
}

func getJsonTag[T any](fieldName string, arg T) string {
	var jsonTag string
	field, ok := reflect.TypeOf(&arg).Elem().FieldByName(fieldName)
	if !ok {
		return fieldName
	}
	fieldJSONName, ok := field.Tag.Lookup("json")
	if !ok {
		return fieldName
	}
	if fieldJSONName != "" {
		jsonTag = fieldJSONName
	} else {
		jsonTag = fieldName
	}
	return jsonTag
}
