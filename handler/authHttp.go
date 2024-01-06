// http/user_handler.go

package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AnuragChaubey2/E-Comm-User-Auth.git/models"
	"github.com/AnuragChaubey2/E-Comm-User-Auth.git/service"
)

type UserHandler struct {
	authService service.AuthService
}

func NewUserHandler(authService service.AuthService) *UserHandler {
	return &UserHandler{authService: authService}
}

// RegistrationHandler handles user registration.
func (uh *UserHandler) RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var registrationRequest models.RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&registrationRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := uh.authService.RegisterUser(
		registrationRequest.Username,
		registrationRequest.Email,
		registrationRequest.Password,
		registrationRequest.Role,
	)
	if err != nil {
		handleRegistrationError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func handleRegistrationError(w http.ResponseWriter, err error) {
	var statusCode int
	switch err {
	case service.ErrInvalidEmail, service.ErrInvalidPassword, service.ErrInvalidRole:
		statusCode = http.StatusBadRequest
	case service.ErrEmailExists:
		statusCode = http.StatusConflict
	default:
		statusCode = http.StatusInternalServerError
	}

	http.Error(w, err.Error(), statusCode)
}