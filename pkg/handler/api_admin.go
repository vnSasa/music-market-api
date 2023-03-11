package handler

import (
	"net/http"
	"strconv"

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
	var input model.SongList
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "main_page_admin.html", "invalid input body")

		return
	}
	err := h.services.Songs.CreateSong(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "main_page_admin.html", nil)
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
	songs, err := h.services.Songs.GetAllSongs()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "get_song.html", gin.H{
		"Songs": songs,
	})
}

func (h *Handler) updateArtist(c *gin.Context) {
	id := c.Param("id")

	c.HTML(http.StatusOK, "update_artist.html", gin.H{
		"id": id,
	})
}

func (h *Handler) saveChanges(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}

	var input model.ArtistList
	if err := c.ShouldBind(&input); err != nil {
		c.HTML(http.StatusBadRequest, "main_page_admin.html", "invalid input body")

		return
	}
	err = h.services.Artists.UpdateArtist(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "main_page_admin.html", nil)
}

