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
		auth.GET("/google/login", h.googleLoginHandler)
		auth.GET("/google/callback", h.googleCallbackHandler)
	}

	myPage := router.Group("/home")
	{
		myPage.GET("/:username", h.userIdentity)
	}

	return router
}
