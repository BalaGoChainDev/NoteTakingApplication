package server

import (
	// ThirdParty libs
	"github.com/gin-gonic/gin"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/handlers"
)

func NewServer() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	// Endpoint for user signup & login
	router.POST("/signup", handlers.SignupHandler)
	router.POST("/login", handlers.LoginHandler)

	// Endpoint for notes
	router.GET("/notes", handlers.GetNotesHandler)
	router.POST("/notes", handlers.CreateNoteHandler)
	router.DELETE("/notes", handlers.DeleteNoteHandler)

	return router
}
