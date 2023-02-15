package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
