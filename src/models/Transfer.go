package models

import "time"

type Transfer struct {
	ID                     uint      `json:"id" gorm:"primaryKey"`
	Account_origin_id      uint      `json:"account_origin_id"`
	Account_destination_id uint      `json:"account_destination_id"`
	Amount                 float64   `json:"amount"`
	Created_at             time.Time `json:"created_at" gorm:"autoCreateTime"`
}
