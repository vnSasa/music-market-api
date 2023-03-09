package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vnSasa/music-market-api/model"
)

func (h *Handler) mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main_page_admin.html", nil)
}

func (h *Handler) createArtist(c *gin.Context) {
	var input model.ArtistList
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "main_page_admin.html", "invalid input body")

		return
	}
	err := h.services.Artists.CreateArtist(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "main_page_admin.html", nil)
}

func (h *Handler) createSong(c *gin.Context) {
}

func (h *Handler) getAllArtist(c *gin.Context) {
	artists, err := h.services.Artists.GetAllArtists()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.HTML(http.StatusOK, "get_artist.html", gin.H{
		"Artists": artists,
	})
}

func (h *Handler) getAllSong(c *gin.Context) {
}
