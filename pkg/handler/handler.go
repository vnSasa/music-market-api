package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vnSasa/music-market-api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoute() *gin.Engine {
	h.initAdmin(nil)
	
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

	admin := router.Group("/api_admin")
	{
		admin.POST("/create_artist", h.createArtist)
		admin.GET("/artist", h.getAllArtist)
		admin.POST("/create_song", h.createSong)
		admin.GET("/song", h.getAllSong)
	}

	return router
}
