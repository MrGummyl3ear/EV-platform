package handler

import (
	"backend/pkg/service"
	"os"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	router := gin.New()

	router.Static("/static", os.Getenv("STATIC_DIR"))

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

	payPage := router.Group("/payment")
	{
		payPage.GET("/config", h.handleConfig)
		payPage.POST("/create-payment-intent", h.createPayment)
		payPage.POST("/webhook", h.handleWebhook)
	}

	chatPage := router.Group("/posts", h.userIdentity)
	{
		chatPage.POST("/create", h.createPost)
		chatPage.GET("/", h.getAllPosts)
		chatPage.POST("/:id/create", h.createComment)
		chatPage.GET("/:id", h.GetAllComments)
	}

	return router
}
