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

	router := gin.Default()

	router.Static("/js", "./pkg/handler/templates/js")
	router.LoadHTMLGlob("./pkg/handler/templates/*.html")

	router.GET("/", h.index)

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", h.signUp)
		auth.POST("/new-sign-up", h.newSignUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/logout", h.logout)
	}

	auth.Use(saveAccessToken)

	admin := router.Group("/api_admin", h.adminIdentity)
	{
		admin.GET("/main_page", h.mainPageAdmin)

		admin.GET("/sign-up_admin", h.signUpAdmin)
		admin.POST("/add_new_admin", h.addNewAdmin)

		admin.GET("/create_artist", h.createArtist)
		admin.POST("/save_artist", h.saveArtist)
		admin.GET("/update_artist/:id", h.updateArtist)
		admin.PUT("/save_changes_artist/:id", h.saveChangesArtist)
		admin.GET("/artist", h.getAllArtist)
		admin.DELETE("/delete_artist/:id", h.deleteArtist)

		admin.GET("/create_song", h.createSong)
		admin.POST("/save_song", h.saveSong)
		admin.GET("/update_song/:id", h.updateSong)
		admin.PUT("/save_changes_song/:id", h.saveChangesSong)
		admin.GET("/song", h.getAllSong)
		admin.GET("/playlist/:id", h.getPlaylist)
		admin.DELETE("/delete_song/:id", h.deleteSong)
	}

	user := router.Group("/api_user", h.userIdentity)
	{
		user.GET("/main_page", h.mainPageUser)

		user.GET("/user_data", h.userData)
		user.PUT("/update_data", h.updateUser)

		user.GET("/get_song", h.getSongs)
		user.GET("/song_info/:id", h.getSongInfo)
		
		user.GET("/get_artist", h.getArtist)
		user.GET("/playlist/:id", h.getPlaylistByArtist)
		
		user.GET("/user_playlist", h.getUserPlaylist)
		user.POST("/add_to_playlist/:id", h.addToPlaylist)
		user.DELETE("/delete_from_playlist/:id", h.deleteSongFromPlaylist)
		
		user.GET("/user_toplist", h.getUserToplist)
		user.POST("/add_to_toplist/:id", h.addToToplist)
		user.DELETE("/delete_from_toplist/:id", h.deleteSongFromToplist)

		user.GET("/create_new_data", h.createNewData)
		user.POST("/save_new_data", h.saveNewData)
		user.GET("/create_new_song", h.createNewSong)
		user.POST("/save_new_song", h.saveNewSong)
	}

	return router
}
