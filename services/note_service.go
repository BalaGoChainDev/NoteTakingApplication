package services

import (
	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/models"
	"github.com/BalaGoChainDev/NoteTakingApplication/repositories"
	"github.com/BalaGoChainDev/NoteTakingApplication/response"
)

// NoteService provides methods to handle note-related operations.
type NoteService interface {
	CreateNote(note *models.Note) (uint, response.ErrorResponder)
	GetNotes(sessionID string) ([]*models.Note, response.ErrorResponder)
	DeleteNote(sID string, noteID string) response.ErrorResponder
}

type noteService struct {
	noteRepo repositories.NoteRepository
}

// NewNoteService creates a new instance of NoteService.
func NewNoteService() NoteService {
	return &noteService{
		noteRepo: repositories.NewNoteRepository(),
	}
}

// CreateNote creates a new note.
func (n *noteService) CreateNote(note *models.Note) (uint, response.ErrorResponder) {
	return n.noteRepo.Create(note)
}

// GetNotes retrieves all notes for a given session ID.
func (n *noteService) GetNotes(sessionID string) ([]*models.Note, response.ErrorResponder) {
	sessionService := repositories.NewSessionRepository()
	session, err := sessionService.GetSession(sessionID)
	if err != nil {
		return nil, err
	}
	return n.noteRepo.GetAllByUserID(session.Email)
}

// DeleteNote deletes a note with the specified note ID, for the given session ID.
func (n *noteService) DeleteNote(sID string, noteID string) response.ErrorResponder {
	sessionService := NewSessionService()
	userInfo, err := sessionService.GetSessionUserInfo(sID)
	if err != nil {
		return err
	}
	return n.noteRepo.DeleteByID(noteID, userInfo.Email)
}
