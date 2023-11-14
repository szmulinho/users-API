package endpoints

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/szmulinho/users/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (h *handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser model.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newUser.Password = string(hashedPassword)

	var publicRepos []model.PublicRepo
	referrer := r.Header.Get("Referer")
	if strings.Contains(referrer, "https://szmul-med-github-login.onrender.com/") {
		hasSzmulMedRepo := false
		for _, repo := range publicRepos {
			if repo.Name == "szmul-med" {
				hasSzmulMedRepo = true
				break
			}
		}

		if hasSzmulMedRepo {
			newUser.Role = "doctor"
		} else {
			newUser.Role = "user"
		}
	} else {
		// Jeśli żądanie nie pochodzi z Twojego GitHub API, przypisz standardową rolę "user"
		newUser.Role = "user"
	}

	result := h.db.Create(&newUser)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	userJSON, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(userJSON)
}
