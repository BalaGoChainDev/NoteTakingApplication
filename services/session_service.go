package services

import (
	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/models"
	"github.com/BalaGoChainDev/NoteTakingApplication/repositories"
	"github.com/BalaGoChainDev/NoteTakingApplication/response"
)

type SessionService interface {
	GetSessionId(emailId string) (string, response.ErrorResponder)
	GetSessionUserInfo(sessionID string) (*models.User, response.ErrorResponder)
}

type sessionService struct {
	SessionRepo repositories.SessionRepository
}

// NewSessionService creates a new instance of SessionService.
func NewSessionService() SessionService {
	return &sessionService{
		SessionRepo: repositories.NewSessionRepository(),
	}
}

// GetSessionID retrieves the session ID for the given email ID.
func (s *sessionService) GetSessionId(emailId string) (string, response.ErrorResponder) {
	return s.SessionRepo.UpdateSessionID(emailId)
}

// GetSessionUserInfo retrieves the user information associated with the given session ID.
func (s *sessionService) GetSessionUserInfo(sessionID string) (*models.User, response.ErrorResponder) {
	var sessionUser *models.User
	sessionInfo, err := s.SessionRepo.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	userRepo := repositories.NewUserRepository()
	sessionUser, err = userRepo.GetByEmail(sessionInfo.Email)
	if err != nil {
		return nil, err
	}

	sessionUser.Session = *sessionInfo
	return sessionUser, nil
}
