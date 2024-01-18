package api

import (
	"github.com/Divyansh12/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validateCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
