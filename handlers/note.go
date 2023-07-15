package handlers

import (
	// Standard libs
	"net/http"

	// ThirdParty libs
	"github.com/gin-gonic/gin"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/models"
	"github.com/BalaGoChainDev/NoteTakingApplication/response"
	"github.com/BalaGoChainDev/NoteTakingApplication/services"
	"github.com/BalaGoChainDev/NoteTakingApplication/utils"
)

// GetRequest represents the request body for the GetNotesHandler.
type GetRequest struct {
	Sid string `json:"sid" validate:"required"`
}

// PostRequest represents the request body for the CreateNoteHandler.
type PostRequest struct {
	Sid  string `json:"sid" validate:"required"`
	Note string `json:"note" validate:"required"`
}

// DeleteRequest represents the request body for the DeleteNoteHandler.
type DeleteRequest struct {
	Sid    string `json:"sid" validate:"required"`
	NoteID string `json:"id" validate:"required"`
}

// GetNotesHandler handles the get notes endpoint.
func GetNotesHandler(c *gin.Context) {
	var req GetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Invalid request format"))
		return
	}

	if err := utils.ValidateInput(&req); err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}

	noteService := services.NewNoteService()
	notes, err := noteService.GetNotes(req.Sid)
	if err != nil {
		c.JSON(err.GetStatusCode(), err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"notes": notes})
}

// CreateNoteHandler handles the create note endpoint.
func CreateNoteHandler(c *gin.Context) {
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Invalid request format"))
		return
	}

	if err := utils.ValidateInput(&req); err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}

	sessionService := services.NewSessionService()
	userInfo, err := sessionService.GetSessionUserInfo(req.Sid)
	if err != nil {
		c.JSON(err.GetStatusCode(), err.Error())
		return
	}

	note := &models.Note{
		Email: userInfo.Email,
		Note:  req.Note,
	}

	noteService := services.NewNoteService()
	id, err := noteService.CreateNote(note)
	if err != nil {
		c.JSON(err.GetStatusCode(), err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// DeleteNoteHandler handles the delete note endpoint.
func DeleteNoteHandler(c *gin.Context) {
	var req DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Invalid request format"))
		return
	}

	if err := utils.ValidateInput(&req); err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}

	noteService := services.NewNoteService()
	err := noteService.DeleteNote(req.Sid, req.NoteID)
	if err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
