package endpoints

import (
	"github.com/szmulinho/users/internal/model"
	"gorm.io/gorm"
	"net/http"
)

type Handlers interface {
	GenerateToken(w http.ResponseWriter, r *http.Request, ID int64, isUser bool) (string, error)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUserDataHandler(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc
	getUserFromToken(tokenString string) (*model.User, error)
}

type handlers struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) Handlers {
	return &handlers{
		db: db,
	}
}
