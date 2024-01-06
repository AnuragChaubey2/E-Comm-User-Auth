// service/auth.go

package service

import (
	"errors"
	"regexp"

	"github.com/AnuragChaubey2/E-Comm-User-Auth.git/store"
	"github.com/AnuragChaubey2/E-Comm-User-Auth.git/models"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidRole     = errors.New("invalid role")
	ErrEmailExists     = errors.New("email already exists")
)

type AuthService struct {
	userStore store.UserStore
}

func NewAuthService(userStore store.UserStore) *AuthService {
	return &AuthService{
		userStore: userStore,
	}
}

func (s *AuthService) ValidateRegistrationRequest(username, email, password, role string) error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return ErrInvalidEmail
	}

	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()-_=+]{6,}$`)
	if !passwordRegex.MatchString(password) {
		return ErrInvalidPassword
	}

	if role != "admin" && role != "user" {
		return ErrInvalidRole
	}

	return nil
}

func (s *AuthService) RegisterUser(username, email, password, role string) (*models.RegistrationResponse, error) {
	if err := s.ValidateRegistrationRequest(username, email, password, role); err != nil {
		return nil, err
	}

	emailExists, err := s.userStore.IsEmailExists(email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return nil, ErrEmailExists
	}

	token, err := GenerateJWTToken(username, role)
	if err != nil {
		return nil, err
	}

	response := &models.RegistrationResponse{
		Status: "Success",
		Token:  token,
	}

	return response, nil
}