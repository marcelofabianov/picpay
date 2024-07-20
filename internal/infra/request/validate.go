package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/marcelofabianov/picpay/internal/infra/response"
)

func Init() *validator.Validate {
	v := validator.New()

	//... add custom validations

	return v
}

func IsValid(c *fiber.Ctx, req interface{}) bool {
	v := Init()

	if err := v.Struct(req); err != nil {
		validationErrors, _ := err.(validator.ValidationErrors)
		var errors []map[string]string
		for _, e := range validationErrors {
			errors = append(errors, map[string]string{
				"field":   e.Field(),
				"message": e.Tag(),
			})
		}
		response.BadRequestErrors(c, errors)

		return false
	}

	return true
}
