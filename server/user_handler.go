package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/go-api-test/input"
	"example.com/go-api-test/service"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := h.userService.GetUsers(ctx)

	if err != nil {
		log.Printf("Error getting users: %v", err)
		http.Error(w, "Error getting users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if payload.Name == "" || payload.Email == "" {
		http.Error(w, "Name and email are required", http.StatusBadRequest)
		return
	}

	input := input.CreateUserInput{
		Name:  payload.Name,
		Email: payload.Email,
	}

	user, err := h.userService.CreateUser(r.Context(), input)

	if err != nil {

		if err == service.ErrEmailAlreadyExists {
			http.Error(w, "Email already in use", http.StatusConflict)
			return
		}

		http.Error(w, "Error creating user", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(r.Context(), id)

	if err != nil {
		if err == service.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var input input.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if input.Name == nil && input.Email == nil {
		http.Error(w, "No fields to update", http.StatusBadRequest)
		return
	}

	user, err := h.userService.UpdateUser(r.Context(), id, input)

	if err != nil {
		if err == service.ErrEmailAlreadyExists {
			http.Error(w, "Email already in use", http.StatusConflict)
			return
		}

		if err == service.ErrUserNotFound {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
