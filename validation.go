package validation

import (
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Validation struct {
	validate *validator.Validate
	Errors   map[string]string
}

func New() *Validation {
	return &Validation{
		validate: validator.New(),
		Errors:   make(map[string]string),
	}
}

func (v *Validation) valid(field interface{}, key string, message string, tag string) bool {
	if err := v.validate.Var(field, tag); err != nil {
		v.Errors[key] = message
		return false
	}
	return true
}

func (v *Validation) ASCII(field interface{}, key string, message string) bool {
	return v.valid(field, key, message, "ascii")
}

func (v *Validation) Email(field interface{}, key string, message string) bool {
	return v.valid(field, key, message, "email")
}

func (v *Validation) Latitude(field interface{}, key string, message string) bool {
	return v.valid(field, key, message, "latitude")
}

func (v *Validation) Longitude(field interface{}, key string, message string) bool {
	return v.valid(field, key, message, "longitude")
}

func (v *Validation) Max(field interface{}, max int, key string, message string) bool {
	return v.valid(field, key, message, "max="+strconv.Itoa(max))
}

func (v *Validation) Min(field interface{}, min int, key string, message string) bool {
	return v.valid(field, key, message, "min="+strconv.Itoa(min))
}

func (v *Validation) Numeric(field interface{}, min int, key string, message string) bool {
	return v.valid(field, key, message, "numeric")
}

func (v *Validation) Required(field interface{}, key string, message string) bool {
	return v.valid(field, key, message, "required")
}

func (v *Validation) URL(field interface{}, key string, message string) bool {
	return v.valid(field, key, message, "url")
}

func (v *Validation) HasErrors() bool {
	return len(v.Errors) > 0
}