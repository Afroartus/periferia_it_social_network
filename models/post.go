package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Message   string    `json:"message"`
	UserID    uint      `json:"user_id"`
	Likes     int       `gorm:"default:0" json:"likes"`
	CreatedAt time.Time `json:"created_at"`
}
