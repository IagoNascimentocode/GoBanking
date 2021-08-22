package models

type Credentials struct {
	Cpf    string `json:"cpf" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}
