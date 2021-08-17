package models

import "time"

type Account struct {
	ID         uint `gorm:"primaryKey"`
	Cpf        int
	Name       string
	Secret     string
	Balance    float64   `goorm:"default:0"`
	Created_at time.Time `gorm:"autoCreateTime"`
}
