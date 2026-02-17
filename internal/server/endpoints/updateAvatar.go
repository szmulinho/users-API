package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/users/internal/model"
)

type UpdateAvatarRequest struct {
	AvatarURL string `json:"avatar_url"`
}

func (h *handlers) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	// Pobierz token z headera
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header required", http.StatusUnauthorized)
		return
	}

	// Parse token
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return model.JwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Pobierz user ID z claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		return
	}

	userID, ok := claims["userID"].(float64)
	if !ok {
		http.Error(w, "Invalid userID in token", http.StatusUnauthorized)
		return
	}

	// Pobierz dane z request body
	var req UpdateAvatarRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Znajdź użytkownika
	var user model.User
	result := h.db.First(&user, int64(userID))
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Zaktualizuj avatar
	user.AvatarURL = req.AvatarURL
	result = h.db.Save(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Nie zwracaj hasła w odpowiedzi
	user.Password = ""

	userJSON, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

