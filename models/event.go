package models

import "time"

type Event struct {
	Id        string    `json:"id" gorm:"primaryKey;autoIncrement"`
	Timestamp time.Time `json:"timestamp" gorm:"autoCreateTime"`
	Type      string    `json:"type" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	UserId    string    `json:"user_id" gorm:"not null"`
}

type CreateEventRequest struct {
	Type    string `json:"type" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateEventRequest struct {
	Type    string `json:"type"`
	UserId  string `json:"user_id"`
	Content string `json:"content"`
}
