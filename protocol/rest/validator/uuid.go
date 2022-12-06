package validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func ValidateUUID(id string) error {
	return validation.Validate(
		id,
		validation.Required,
		is.UUIDv4,
	)
}
