package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/users/internal/model"
	"net/http"
)

func (h *handlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var Users []model.User
	if err := h.db.Find(&Users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Users)
}
