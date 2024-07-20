package validate

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/marcelofabianov/picpay/internal/adapter/api/common/response"
)

func Init() *validator.Validate {
	v := validator.New()

	//v.RegisterValidation("cpf", ValidateCPF)
	//v.RegisterValidation("cnpj", ValidateCNPJ)
	//v.RegisterValidation("cnpj_cpf", ValidateDocumentRegistry)

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
		response.BadRequestErrorsResponse(c, errors)

		return false
	}

	return true
}

func ValidateDocumentRegistry(fl validator.FieldLevel) bool {
	documentRegistry := fl.Field().String()
	documentRegistry = strings.Join(regexp.MustCompile(`\D`).FindAllString(documentRegistry, -1), "")

	if len(documentRegistry) != 14 && len(documentRegistry) != 11 {
		return false
	}

	if len(documentRegistry) == 14 {
		return ValidateCNPJ(fl)
	}

	return ValidateCPF(fl)
}

func ValidateCNPJ(fl validator.FieldLevel) bool {
	cnpj := fl.Field().String()
	cnpj = strings.Join(regexp.MustCompile(`\D`).FindAllString(cnpj, -1), "")

	if len(cnpj) != 14 {
		return false
	}

	//...

	return true
}

func ValidateCPF(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()
	cpf = strings.Join(regexp.MustCompile(`\D`).FindAllString(cpf, -1), "")

	if len(cpf) != 11 {
		return false
	}

	//...

	return true
}
