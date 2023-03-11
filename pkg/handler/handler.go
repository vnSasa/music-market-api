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
	h.initAdmin()

	router := gin.New()

	router.Static("/js", "./pkg/handler/templates/js")
	router.LoadHTMLGlob("./pkg/handler/templates/*.html")

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
		admin.GET("/main_page", h.mainPage)
		admin.POST("/create_artist", h.createArtist)
		admin.GET("/update_artist/:id", h.updateArtist)
		admin.POST("/save_changes/:id", h.saveChanges)
		admin.GET("/artist", h.getAllArtist)
		admin.POST("/create_song", h.createSong)
		admin.GET("/song", h.getAllSong)
		// admin.POST("/update_song", h.updateSong)
	}

	return router
}
