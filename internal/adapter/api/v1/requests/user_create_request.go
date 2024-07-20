package requests

type UserCreateRequest struct {
	Name             string `json:"name" validate:"required,min=3,max=100"`
	Email            string `json:"email" validate:"required,email,min=3,max=150"`
	Password         string `json:"password" validate:"required,min=6,max=100"`
	DocumentRegistry string `json:"document_registry" validate:"required,min=11,max=14,cnpj_cpf"`
}
