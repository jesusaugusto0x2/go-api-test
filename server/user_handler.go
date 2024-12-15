package server

import (
	"encoding/json"
	"fmt"
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
		msg, status := buildErrorResponse(err)
		http.Error(w, msg, status)
		return
	}

	respondJSON(w, http.StatusOK, users)
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, ErrInvalidPayloadMsg, http.StatusBadRequest)
		return
	}

	if payload.Name == "" || payload.Email == "" {
		http.Error(w, ErrNoFieldsToUpdateMsg, http.StatusBadRequest)
		return
	}

	input := input.CreateUserInput{
		Name:  payload.Name,
		Email: payload.Email,
	}

	user, err := h.userService.CreateUser(r.Context(), input)

	if err != nil {
		msg, status := buildErrorResponse(err)
		http.Error(w, msg, status)
		return
	}

	respondJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)

	if err != nil {
		http.Error(w, ErrInvalidUserIDMsg, http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUser(r.Context(), id)

	if err != nil {
		msg, status := buildErrorResponse(err)
		http.Error(w, msg, status)
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)

	if err != nil {
		http.Error(w, ErrInvalidUserIDMsg, http.StatusBadRequest)
		return
	}

	var input input.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, ErrInvalidPayloadMsg, http.StatusBadRequest)
		return
	}

	if input.Name == nil && input.Email == nil {
		http.Error(w, ErrInvalidPayloadMsg, http.StatusBadRequest)
		return
	}

	user, err := h.userService.UpdateUser(r.Context(), id, input)

	if err != nil {
		msg, status := buildErrorResponse(err)
		http.Error(w, msg, status)
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r)

	if err != nil {
		http.Error(w, ErrInvalidUserIDMsg, http.StatusBadRequest)
		return
	}

	err = h.userService.DeleteUser(r.Context(), id)

	if err != nil {
		msg, status := buildErrorResponse(err)
		http.Error(w, msg, status)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Private functions
func parseIDParam(r *http.Request) (int, error) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return -1, fmt.Errorf("invalid user id param")
	}

	return id, nil
}

func buildErrorResponse(err error) (string, int) {
	switch err {
	case service.ErrEmailAlreadyExists:
		return "Email already in use", http.StatusConflict
	case service.ErrUserNotFound:
		return "User not found", http.StatusNotFound
	default:
		return "Internal server error", http.StatusInternalServerError
	}
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
