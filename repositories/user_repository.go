package repositories

import (
	// Standard libs
	"fmt"
	"net/http"
	"strings"

	// ThirdParty libs
	"gorm.io/gorm"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/models"
	"github.com/BalaGoChainDev/NoteTakingApplication/response"
)

// UserRepository defines the methods for working with users in the database.
type UserRepository interface {
	Create(user *models.User) response.ErrorResponder
	GetByEmail(email string) (*models.User, response.ErrorResponder)
}

type userRepo struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository() UserRepository {
	return &userRepo{
		db: GetDB(),
	}
}

// Create creates a new user in the database.
func (u *userRepo) Create(user *models.User) response.ErrorResponder {
	err := u.db.Create(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return response.NewErrorResponse(409, "Email conflict!")
		}
		return response.NewErrorResponse(500, fmt.Sprintf("Failed to register user: %v", err.Error()))
	}

	return nil
}

// GetByEmail retrieves a user from the database based on the email.
func (u *userRepo) GetByEmail(email string) (*models.User, response.ErrorResponder) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, response.NewErrorResponse(http.StatusUnauthorized, "Unauthorized user! Email does not exisits.")
		}
		return nil, response.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to login: %v", err.Error()))
	}
	return &user, nil
}
