package services

import (
	// Standard libs
	"net/http"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/models"
	"github.com/BalaGoChainDev/NoteTakingApplication/repositories"
	"github.com/BalaGoChainDev/NoteTakingApplication/response"
)

// UserService provides methods to handle user-related operations.
type UserService interface {
	Register(user *models.User) response.ErrorResponder
	Login(credentials *models.Credentials) (*models.User, response.ErrorResponder)
}

type userService struct {
	userRepo repositories.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService() UserService {
	return &userService{
		userRepo: repositories.NewUserRepository(),
	}
}

// Register creates a new user.
func (us *userService) Register(user *models.User) response.ErrorResponder {
	err := us.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

// Login authenticates the user based on the provided credentials.
func (us *userService) Login(credentials *models.Credentials) (*models.User, response.ErrorResponder) {
	user, err := us.userRepo.GetByEmail(credentials.Email)
	if err != nil {
		return nil, err
	}

	if user.Password != credentials.Password {
		return nil, response.NewErrorResponse(http.StatusUnauthorized, "Unauthorized user! Incorrect password.")
	}

	return user, nil
}
