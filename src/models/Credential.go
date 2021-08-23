package models

type Credential struct {
	Cpf    string `json:"cpf" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}
