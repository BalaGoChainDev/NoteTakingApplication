package repositories

import (
	// Standard libs
	"errors"
	"fmt"
	"net/http"

	// ThirdParty libs
	"gorm.io/gorm"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/models"
	"github.com/BalaGoChainDev/NoteTakingApplication/response"
)

// NoteRepository defines the methods for working with notes in the database.
type NoteRepository interface {
	Create(note *models.Note) (uint, response.ErrorResponder)
	GetAllByUserID(emailId string) ([]*models.Note, response.ErrorResponder)
	DeleteByID(noteID string, emailId string) response.ErrorResponder
}

type noteRepo struct {
	db *gorm.DB
}

// NewNoteRepository creates a new instance of the NoteRepository.
func NewNoteRepository() NoteRepository {
	return &noteRepo{
		db: GetDB(),
	}
}

// Create creates a new note in the database.
func (n *noteRepo) Create(note *models.Note) (uint, response.ErrorResponder) {
	err := n.db.Create(note).Error
	if err != nil {
		return 0, response.NewErrorResponse(500, fmt.Sprintf("Failed to create note: %v", err.Error()))
	}
	return note.ID, nil
}

// GetAllByUserID retrieves all notes associated with a user by their email ID.
func (n *noteRepo) GetAllByUserID(emailId string) ([]*models.Note, response.ErrorResponder) {
	var notes []*models.Note
	err := n.db.Where("email = ?", emailId).Find(&notes).Error
	if err != nil {
		return nil, response.NewErrorResponse(500, fmt.Sprintf("Failed to get notes: %v", err.Error()))
	}
	return notes, nil
}

// DeleteByID deletes a note by its ID, checking if it belongs to the specified email ID.
func (n *noteRepo) DeleteByID(noteID string, emailID string) response.ErrorResponder {
	// Check if the note exists
	var note models.Note
	result := n.db.Where("id = ?", noteID).First(&note)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return response.NewErrorResponse(http.StatusBadRequest, "Bad request! Note ID does not exist")
		}
		return response.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to delete note: %v", result.Error))
	}

	if note.Email != emailID {
		return response.NewErrorResponse(http.StatusUnauthorized, "Unauthorized user can't able to delete the note")
	}

	// Delete the note
	result = n.db.Delete(&note)
	if result.Error != nil {
		return response.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("Failed to delete note: %v", result.Error))
	}
	if result.RowsAffected == 0 {
		return response.NewErrorResponse(http.StatusBadRequest, "Bad request! Note ID does not exist")
	}

	return nil
}
