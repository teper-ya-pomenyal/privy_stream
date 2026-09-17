package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/teper-ya-pomenyal/privy_stream/gateway/internal/clients"
)

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type UserHandler struct {
	userClient *clients.UserClient
}

func NewUserHandler(userClient *clients.UserClient) *UserHandler {
	return &UserHandler{userClient: userClient}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.UserName == "" || req.Password == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	res, err := h.userClient.Login(r.Context(), req.UserName, req.Password)
	if err != nil {
		mapGRPCError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.BirthDate == "" || req.Password == "" || req.UserName == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	res, err := h.userClient.Register(r.Context(), req.UserName, req.Password, birthDate)
	if err != nil {
		mapGRPCError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.RefreshToken == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	res, err := h.userClient.Refresh(r.Context(), req.RefreshToken)
	if err != nil {
		mapGRPCError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.RefreshToken == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.userClient.Logout(r.Context(), req.RefreshToken)
	if err != nil {
		mapGRPCError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)

}
