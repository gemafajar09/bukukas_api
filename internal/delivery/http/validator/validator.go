package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func PesanError(err error) map[string]string {
	errorsMap := make(map[string]string)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			field := fieldError.Field()

			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorsMap[field] = "Invalid email format"
			case "unique":
				errorsMap[field] = fmt.Sprintf("%s already exists", field)
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field)
			case "gte":
				errorsMap[field] = fmt.Sprintf("%s must be greater than or equal to %s", field, fieldError.Param())
			case "lte":
				errorsMap[field] = fmt.Sprintf("%s must be less than or equal to %s", field, fieldError.Param())
			case "gt":
				errorsMap[field] = fmt.Sprintf("%s must be greater than %s", field, fieldError.Param())
			case "lt":
				errorsMap[field] = fmt.Sprintf("%s must be less than %s", field, fieldError.Param())
			case "date":
				errorsMap[field] = fmt.Sprintf("%s must be a valid date", field)
			case "time":
				errorsMap[field] = fmt.Sprintf("%s must be a valid time", field)
			case "datetime":
				errorsMap[field] = fmt.Sprintf("%s must be a valid datetime", field)
			case "url":
				errorsMap[field] = fmt.Sprintf("%s must be a valid URL", field)
			case "uuid":
				errorsMap[field] = fmt.Sprintf("%s must be a valid UUID", field)
			case "phone":
				errorsMap[field] = fmt.Sprintf("%s must be a valid phone number", field)
			case "json":
				errorsMap[field] = fmt.Sprintf("%s must be a valid JSON", field)
			default:
				errorsMap[field] = "Invalid value"
			}
		}
	}

	if err != nil {
		if JikaDuplikast(err) {
			switch {
			case strings.Contains(err.Error(), "username"):
				errorsMap["username"] = "Username already exists"
			case strings.Contains(err.Error(), "email"):
				errorsMap["email"] = "Email already exists"
			case strings.Contains(err.Error(), "phone"):
				errorsMap["phone"] = "Phone number already exists"
			case strings.Contains(err.Error(), "uuid"):
				errorsMap["uuid"] = "UUID already exists"
			default:
				errorsMap["duplicate"] = "Duplicate entry"
			}
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			errorsMap["error"] = "Record not found"
		}
	}

	return errorsMap
}

func JikaDuplikast(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}
