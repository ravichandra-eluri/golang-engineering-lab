package handler

import (
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	Secret string
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.Email == "" || req.Password == "" {
		http.Error(w, "email and password required", http.StatusBadRequest)
		return
	}
	// TODO: validate credentials and return a signed JWT
	w.WriteHeader(http.StatusOK)
}
