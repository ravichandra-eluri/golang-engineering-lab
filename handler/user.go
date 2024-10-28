package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"golang-engineering-lab/service"
)

type UserHandler struct {
	Users *service.UserService
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.Users.List(r.Context())
	if err != nil {
		slog.Error("list users", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
