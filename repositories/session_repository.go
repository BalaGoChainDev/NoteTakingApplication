package repositories

import (
	// Standard libs
	"fmt"
	"net/http"

	// ThirdParty libs
	"gorm.io/gorm"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/models"
	"github.com/BalaGoChainDev/NoteTakingApplication/response"
	"github.com/google/uuid"
)

// SessionRepository defines the methods for working with sessions in the database.
type SessionRepository interface {
	UpdateSessionID(emailId string) (string, response.ErrorResponder)
	GetSession(sessionID string) (*models.Session, response.ErrorResponder)
}

type sessionRepo struct {
	db *gorm.DB
}

// NewSessionRepository creates a new instance of SessionRepository.
func NewSessionRepository() SessionRepository {
	return &sessionRepo{
		db: GetDB(),
	}
}

// UpdateSessionID updates the session ID for the specified email ID.
func (s *sessionRepo) UpdateSessionID(emailId string) (string, response.ErrorResponder) {
	sessionID := uuid.NewString()
	session := models.Session{Email: emailId, SessionID: sessionID}

	err := s.db.Model(&session).Save(&session).Error
	if err != nil {
		return "", response.NewErrorResponse(500, fmt.Sprintf("Failed to update session id: %v", err.Error()))
	}
	return sessionID, nil
}

// GetSession retrieves the session information for the specified session ID.
func (s *sessionRepo) GetSession(sessionID string) (*models.Session, response.ErrorResponder) {
	var session models.Session
	err := s.db.Where("session_id = ?", sessionID).First(&session).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, response.NewErrorResponse(http.StatusUnauthorized, "Unauthorized user! sid does not exisits.")
		}
		return nil, response.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to get session info: %v", err.Error()))
	}

	return &session, nil
}
