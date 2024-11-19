package handler

import (
	"backend/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

//Have to use Gin mode somehow

func (h *Handler) InitRoutes(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/:provider/login", h.signInWithProvider)
		auth.GET("/:provider/callback", h.callbackHandler)
	}

	myPage := router.Group("/home")
	{
		myPage.GET("/:username", h.userIdentity)
	}

	return router
}
