package model

import (
	"os"
)

var JwtKey = []byte(os.Getenv("JWT_KEY"))

type Exception struct {
	Message string `json:"message"`
}

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique" json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
