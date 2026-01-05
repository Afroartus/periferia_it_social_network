package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Alias     string    `gorm:"unique" json:"alias"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
}
