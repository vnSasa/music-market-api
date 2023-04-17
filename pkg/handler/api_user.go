package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/vnSasa/music-market-api/model"
)

func (h *Handler) mainPageUser(c *gin.Context) {
	c.HTML(http.StatusOK, "main_page.html", nil)
}

func (h *Handler) userData(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	userData, err := h.services.Authorization.GetUserByID(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "user_data.html", gin.H{
		"UserData": userData,
	})
}

func (h *Handler) updateUser(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		c.Redirect(http.StatusBadRequest, "/api_user/user_data")

		return
	}
	var input model.User
	if err := c.ShouldBind(&input); err != nil {
		c.Redirect(http.StatusBadRequest, "/api_user/user_data")

		return
	}
	err = h.services.Authorization.UpdateUser(userID, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
}

func (h *Handler) getSongs(c *gin.Context) {
	var songList []model.SongData
	songs, err := h.services.Songs.GetAllSongs()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	artists, err := h.services.Artists.GetAllArtists()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	artistMap := make(map[int]string)
	for _, artist := range artists {
		artistMap[artist.ID] = artist.Name
	}
	for _, song := range songs {
		songList = append(songList, model.SongData{
			ID:         song.ID,
			ArtistID:   song.ArtistID,
			ArtistData: artistMap[song.ArtistID],
			Name:       song.Name,
			Genre:      song.Genre,
			Genre2:     song.Genre2,
			Year:       song.Year,
		})
	}
	c.HTML(http.StatusOK, "get_song.html", gin.H{
		"Songs": songList,
	})
}

func (h *Handler) getArtist(c *gin.Context) {
	artists, err := h.services.Artists.GetAllArtists()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "get_artist.html", gin.H{
		"Artists": artists,
	})
}

func (h *Handler) getUserPlaylist(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	var songList []model.SongData
	songs, err := h.services.UsersLibrary.GetUserPlaylist(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	artists, err := h.services.Artists.GetAllArtists()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	artistMap := make(map[int]string)
	for _, artist := range artists {
		artistMap[artist.ID] = artist.Name
	}
	for _, song := range songs {
		songList = append(songList, model.SongData{
			ID:         song.ID,
			ArtistID:   song.ArtistID,
			ArtistData: artistMap[song.ArtistID],
			Name:       song.Name,
			Genre:      song.Genre,
			Genre2:     song.Genre2,
			Year:       song.Year,
		})
	}
	c.HTML(http.StatusOK, "user_playlist.html", gin.H{
		"Songs": songList,
	})
}

func (h *Handler) getPlaylistByArtist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}
	songs, err := h.services.Songs.GetPlaylist(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.HTML(http.StatusOK, "playlist.html", gin.H{
		"Songs": songs,
	})
}

func (h *Handler) addToPlaylist(c *gin.Context) {
	userID, err := h.getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}
	err = h.services.UsersLibrary.AddToPlaylist(userID, songID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.Redirect(http.StatusSeeOther, "/api_user/user_playlist")
}

func (h *Handler) deleteSongFromPlaylist(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}
	err = h.services.UsersLibrary.DeleteSongFromPlaylist(songID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	c.Redirect(http.StatusSeeOther, "/api_user/user_playlist")
}
