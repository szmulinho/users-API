package model

import (
	"gorm.io/gorm"
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

type PublicRepo struct {
	gorm.Model
	GitHubLoginID uint   `gorm:"foreignKey"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}
