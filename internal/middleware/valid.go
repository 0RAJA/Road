package middleware

import (
	"github.com/0RAJA/Road/internal/pkg/times"
	"github.com/go-playground/validator/v10"
	"time"
)

var TimeFormat validator.Func = func(fl validator.FieldLevel) bool {
	if str, ok := fl.Field().Interface().(string); ok {
		if _, err := time.Parse(times.LayoutDateTime, str); err == nil {
			return true
		}
	}
	return false
}
