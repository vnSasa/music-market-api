package handler

import (
	"github.com/vnSasa/music-market-api/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.New()

	router.Static("/static", "static")
	router.LoadHTMLGlob("./pkg/handler/templates/*")
	
	router.GET("/", h.index)

	auth := router.Group("/auth") 
	{
		auth.GET("/sign-up", h.signUp)
		auth.POST("/new-sign-up", h.newSignUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/logout", h.logout)
	}

	auth.Use(h.saveAccessToken)

	return router
}