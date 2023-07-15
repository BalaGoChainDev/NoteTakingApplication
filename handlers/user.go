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

// signupRequest represents the request body for the signup endpoint.
type signupRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

// loginRequest represents the request body for the login endpoint.
type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=50"`
}

// SignupHandler handles the signup endpoint.
func SignupHandler(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Invalid request format"))
		return
	}

	if err := utils.ValidateInput(&req); err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	userService := services.NewUserService()
	if err := userService.Register(user); err != nil {
		c.JSON(err.GetStatusCode(), err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginHandler handles the login endpoint.
func LoginHandler(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse(http.StatusBadRequest, "Invalid request format"))
		return
	}

	if err := utils.ValidateInput(&req); err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}

	credentials := &models.Credentials{
		Email:    req.Email,
		Password: req.Password,
	}

	userService := services.NewUserService()
	_, err := userService.Login(credentials)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	sessionService := services.NewSessionService()
	sessionId, err := sessionService.GetSessionId(credentials.Email)
	if err != nil {
		c.JSON(err.GetStatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "sid": sessionId})
}
