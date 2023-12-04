package models

import "time"

type Event struct {
	Id        string    `json:"id" gorm:"primaryKey;autoIncrement"`
	Timestamp time.Time `json:"timestamp" gorm:"autoCreateTime"`
	Type      string    `json:"type" gorm:"not null"`
	UserId    string    `json:"user_id" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
}
