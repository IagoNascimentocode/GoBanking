package models

type InputAccount struct {
	Name   string `binding:"required"`
	Cpf    int    `binding:"required"`
	Secret string `binding:"required"`
}
