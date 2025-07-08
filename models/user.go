package models

import "time"

type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Name      string
	Email     string
	Role      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
