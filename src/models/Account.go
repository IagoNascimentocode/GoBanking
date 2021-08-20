package models

import "time"

type Account struct {
	ID         uint   `gorm:"primaryKey"`
	Cpf        string `gorm:"unique"`
	Name       string
	Secret     string
	Balance    float64   `goorm:"default:0"`
	Created_at time.Time `gorm:"autoCreateTime"`
}

type Credentials struct {
	Cpf    string `binding:"required"`
	Secret string `binding:"required"`
}
