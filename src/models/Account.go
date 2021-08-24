package models

import "time"

type Account struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Cpf        string    `json:"cpf" gorm:"unique" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Secret     string    `json:"secret" binding:"required"`
	Token      string    `json:"token"`
	Balance    float64   `goorm:"default:0"`
	Created_at time.Time `gorm:"autoCreateTime"`
}
