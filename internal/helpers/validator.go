package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

func ValidateEmail(s string) error {
	err := validate.Var(s, "required,email")
	if err != nil {
		return fmt.Errorf("[email:%s] email is not valid", s)
	}
	return nil
}
